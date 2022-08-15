// Package dizzy contains a implementation of a websocket client to discord
package dizzy

import (
	"compress/zlib"
	"context"
	"crypto/rand"
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	mathrand "math/rand"
	"net"
	"net/http"
	"runtime"
	"sync"
	"time"

	"github.com/SomethingBot/dizzy/libinfo"
	"github.com/SomethingBot/dizzy/logging"
	"github.com/SomethingBot/dizzy/primitives"
	"github.com/SomethingBot/multierror"
	"github.com/gorilla/websocket"
)

type Client struct {
	apikey     string
	gatewayURL string
	intents    primitives.GatewayIntent

	conn      *websocket.Conn
	sequence  syncedCounter
	writeLock sync.Mutex

	eDist  eDist
	logger logging.Logger

	closeClient     chan error
	heartbeatError  chan error
	heartbeatClose  chan struct{}
	readWorkerError chan error
}

func NewClient(apikey, endpoint string, intents primitives.GatewayIntent, eDist eDist, logger logging.Logger) *Client {
	return &Client{
		apikey:          apikey,
		gatewayURL:      endpoint,
		intents:         intents,
		eDist:           eDist,
		logger:          logger,
		closeClient:     make(chan error),
		heartbeatError:  make(chan error),
		heartbeatClose:  make(chan struct{}),
		readWorkerError: make(chan error),
	}
}

func (c *Client) writeToWebsocket(messageType int, data []byte) error {
	c.writeLock.Lock()
	defer c.writeLock.Unlock()
	return c.conn.WriteMessage(messageType, data)
}

//todo: handle resumes

func (c *Client) startWebsocketReader() {
	var (
		err        error
		readCloser io.ReadCloser
	)
	defer func() {
		if readCloser != nil {
			err2 := readCloser.Close()
			if err2 != nil {
				err = fmt.Errorf("could not close readCloser (%v) after error (%w)", err2, err)
			}
		}
		if errors.Is(err, websocket.ErrCloseSent) {
			err = nil
		}
		c.readWorkerError <- err
		close(c.readWorkerError)
	}()

	var (
		messageType  int
		reader       io.Reader
		gatewayEvent primitives.GatewayEvent
		gEvent       primitives.GEvent
		data         []byte
	)
	for {
		messageType, reader, err = c.conn.NextReader()
		if err != nil {
			if errors.Is(err, net.ErrClosed) {
				err = nil
			}
			return
		}

		if messageType == websocket.BinaryMessage {
			if readCloser == nil {
				readCloser, err = zlib.NewReader(readCloser)
				if err != nil {
					return
				}
			}

			err = readCloser.(zlib.Resetter).Reset(reader, nil)
			if err != nil {
				return
			}

			reader = readCloser
		}

		data, err = io.ReadAll(reader)
		if err != nil {
			return
		}

		gEvent = primitives.GEvent{}
		err = json.Unmarshal(data, &gEvent)
		if err != nil {
			if errors.Is(err, &(json.UnmarshalTypeError{})) {
				c.logger.Log(logging.Warning, "gEvent JSON Data does not unmarshal ("+err.Error()+")")
				err = nil
				continue
			} else {
				return
			}
		}

		c.logger.Log(logging.Debug, "GatewayEvent received ("+gEvent.Name+")")

		c.sequence.set(gEvent.SequenceNumber)

		gatewayEvent, err = primitives.GetGatewayEventByName(gEvent.Name)
		if err != nil {
			c.logger.Log(logging.Warning, fmt.Sprintf("GatewayEvent from Discord not found in primitives.GetGatewayEventByName gEvent %#+v", gEvent))
			continue
		}

		err = json.Unmarshal(gEvent.Data, &gatewayEvent)
		if err != nil {
			c.logger.Log(logging.Warning, "gateway event JSON Data does not unmarshal ("+err.Error()+") data ("+string(gEvent.Data)+")")
			err = nil
			continue
		}

		c.eDist.FireEvent(gatewayEvent)
	}

}

func (c *Client) handshake() error {
	messageType, reader, err := c.conn.NextReader()
	if err != nil {
		return fmt.Errorf("could not get next reader (%w)", err)
	}

	var readCloser io.ReadCloser
	if messageType == websocket.BinaryMessage {
		readCloser, err = zlib.NewReader(reader)
		if err != nil {
			return fmt.Errorf("zlib reader could not be created (%w)", err)
		}
		reader = readCloser
		defer func() {
			err2 := readCloser.Close()
			if err2 != nil {
				if err != nil {
					err = fmt.Errorf("could not close zlibReader with error (%v) after error (%w)", err2, err)
				} else {
					err = err2
				}
			}
		}()
	}

	decoder := json.NewDecoder(reader)

	var gEvent primitives.GEvent
	err = decoder.Decode(&gEvent)
	if err != nil {
		return fmt.Errorf("could not decode json gEvent (%w)", err)
	}

	c.sequence.set(gEvent.SequenceNumber)

	//todo: find a way to not have to decode json, store event, then decode json *again*

	var hello primitives.GatewayEventHello
	err = json.Unmarshal(gEvent.Data, &hello)
	if err != nil {
		return fmt.Errorf("could not decode json GatewayEventHello")
	}
	c.eDist.FireEvent(hello)

	var data []byte
	data, err = json.Marshal(primitives.GatewayIdentify{
		Opcode: primitives.GatewayOpcodeIdentify,
		Data: primitives.GatewayIdentifyData{
			Token:   c.apikey,
			Intents: c.intents,
			Properties: primitives.GatewayIdentifyProperties{
				OS:      runtime.GOOS,
				Browser: "dizzy",
				Device:  "dizzy",
			},
		},
	})
	if err != nil {
		return err
	}
	err = c.writeToWebsocket(websocket.TextMessage, data)
	if err != nil {
		return err
	}
	return nil
}

var ErrorNoACKAfterHeartbeat = fmt.Errorf("dizzy: did not receive an ACK after sending a heartbeat")

func (c *Client) generateJitter() float64 { //todo: doesn't need to be on Client struct
	b := make([]byte, 8)
	_, err := rand.Read(b)
	if err != nil {
		c.logger.Log(logging.Warning, "Could not access secure rand for generating jitter, falling back to insecure math/rand with current Unix time as seed. Error ("+err.Error()+")")
		/* #nosec G404 */
		return mathrand.New(mathrand.NewSource(time.Now().Unix())).Float64()
	}
	/* #nosec G404 */
	return mathrand.New(mathrand.NewSource(int64(binary.BigEndian.Uint64(b)))).Float64()
}

func (c *Client) startHeartBeatWorker() {
	request := make(chan struct{})
	c.eDist.RegisterHandler(primitives.GatewayEventTypeHeartbeatRequest, func(event primitives.GatewayEvent) {
		request <- struct{}{}
	})
	ack := make(chan primitives.GatewayEventHeartbeatACK)
	c.eDist.RegisterHandler(primitives.GatewayEventTypeHeartbeatACK, func(event primitives.GatewayEvent) {
		ack <- event.(primitives.GatewayEventHeartbeatACK)
	})
	intervalChange := make(chan int)
	c.eDist.RegisterHandler(primitives.GatewayEventTypeHello, func(event primitives.GatewayEvent) {
		intervalChange <- int(float64(event.(primitives.GatewayEventHello).Interval) * c.generateJitter())
	})

	go func() {
		interval := <-intervalChange //todo: handle a resume, which shouldn't require calling this again

		intervalDuration := time.Duration(interval) * time.Millisecond
		c.logger.Log(logging.Debug, fmt.Sprintf("starting heartbeat-er with interval (%v)", intervalDuration))

		timer := time.NewTimer(intervalDuration)
		hasACKed := true

		var err error

		defer func() {
			if !timer.Stop() {
				<-timer.C
			}

			if errors.Is(err, websocket.ErrCloseSent) {
				err = nil
			}

			c.heartbeatError <- err
			close(c.heartbeatError)
		}()

		for {
			select {
			case <-timer.C:
				if !hasACKed {
					err = ErrorNoACKAfterHeartbeat
					return
				}
				err = c.writeToWebsocket(websocket.TextMessage, []byte(fmt.Sprintf("{\"op\": 1, \"d\":%v}", c.sequence.count())))
				if err != nil {
					return
				}
				timer.Reset(intervalDuration)
			case <-ack:
				hasACKed = true
			case <-request:
				if !timer.Stop() {
					<-timer.C
				}
				hasACKed = false
				err = c.writeToWebsocket(websocket.TextMessage, []byte(fmt.Sprintf("{\"op\": 1, \"d\":%v}", c.sequence.count())))
				if err != nil {
					return
				}
				timer.Reset(intervalDuration)
			case <-c.heartbeatClose:
				return
			}
		}
	}()
}

func (c *Client) startLifecycleWorker() {
	var err error
	var err2 error
	select {
	case <-c.closeClient:
		break
	case err2 = <-c.heartbeatError:
		if err2 != nil {
			err = multierror.Append(err, fmt.Errorf("heartbeatworker error (%w)", err2))
		}
	case err2 = <-c.readWorkerError:
		if err2 != nil {
			err = multierror.Append(err, fmt.Errorf("readWorker error (%w)", err2))
		}
	}

	close(c.heartbeatClose)

	err2 = c.writeToWebsocket(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
	if err2 != nil {
		err = multierror.Append(err, fmt.Errorf("could not write to websocket (%w)", err2))
	}
	err2 = c.conn.Close()
	if err2 != nil {
		err = multierror.Append(err, fmt.Errorf("could not close conn (%w)", err2))
	}

	err2, ok := <-c.heartbeatError
	if ok && err2 != nil {
		err = multierror.Append(err, fmt.Errorf("heartbeatworker error (%w)", err2))
	}
	err2, ok = <-c.readWorkerError
	if ok && err2 != nil {
		err = multierror.Append(err, fmt.Errorf("readWorker error (%w)", err2))
	}

	c.eDist.FireEvent(primitives.GatewayEventClientShutdown{Err: err})
	c.eDist.WaitTilDone()
	c.closeClient <- err
}

// Open Client and return if error while opening and/or setting up a session with Discord
func (c *Client) Open() error {
	c.logger.Log(logging.Info, "Starting Discord Client Library")
	dialer := websocket.Dialer{
		Proxy:            http.ProxyFromEnvironment,
		HandshakeTimeout: time.Second, //todo: make configurable
	}

	gatewayHTTPHeader := http.Header{}
	gatewayHTTPHeader.Set("User-Agent", libinfo.BotUserAgent)
	gatewayHTTPHeader.Set("accept-encoding", "zlib") //todo: make configurable

	var err error
	c.conn, _, err = dialer.DialContext(context.Background(), c.gatewayURL+"/?v=9&encoding=json", gatewayHTTPHeader)
	if err != nil {
		return fmt.Errorf("could not dial (%v) error (%v)", c.gatewayURL, err)
	}
	defer func() {
		if err != nil {
			err2 := c.conn.Close()
			if err2 != nil {
				err = fmt.Errorf("%w, also could not close c.Conn %v", err, err2)
			}
		}
	}()

	c.startHeartBeatWorker()

	err = c.handshake()
	if err != nil {
		return fmt.Errorf("could not handshake (%w)", err)
	}

	go c.startWebsocketReader()

	go c.startLifecycleWorker()

	c.logger.Log(logging.Info, "Client library running")
	return nil
}

// Close client and return if there was a error while closing
func (c *Client) Close() error {
	c.closeClient <- nil
	return <-c.closeClient
}

func (c *Client) AddHandlerFunc(eventType primitives.GatewayEventType, handlerFunc func(event primitives.GatewayEvent)) error {
	c.eDist.RegisterHandler(eventType, handlerFunc)
	return nil
}
