package discordwebapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/SomethingBot/dizzy/libinfo"
)

// SessionStartLimit is described at https://discord.com/developers/docs/topics/gateway#session-start-limit-object
type SessionStartLimit struct {
	Total          int `json:"total"`
	Remaining      int `json:"remaining"`
	ResetAfter     int `json:"reset_after"`
	MaxConcurrency int `json:"max_concurrency"`
}

// GatewayWebsocketInformation contains information for a sharded bot
type GatewayWebsocketInformation struct {
	URL               string            `json:"url"`
	Shards            int               `json:"shards"`
	SessionStartLimit SessionStartLimit `json:"session_start_limit"`
}

// GetGatewayWebsocketInformation gets information to connect to the discord websocket based on the provided apiKey
func GetGatewayWebsocketInformation(discordAPIGatewayURL string, apiKey string) (GatewayWebsocketInformation, error) {
	if discordAPIGatewayURL == "" {
		discordAPIGatewayURL = "https://discord.com/api/gateway/bot"
	}

	req, err := http.NewRequest("GET", discordAPIGatewayURL, nil)
	if err != nil {
		return GatewayWebsocketInformation{}, err
	}
	req.Header.Set("User-Agent", libinfo.BotUserAgent)
	req.Header.Add("Authorization", "Bot "+apiKey)

	resp, err := (&http.Client{}).Do(req)
	if err != nil {
		return GatewayWebsocketInformation{}, err
	}

	defer func() {
		if err != nil {
			err2 := resp.Body.Close()
			if err2 != nil {
				err = fmt.Errorf("could not close body (%v) after error (%w)", err2, err)
			}
		}
	}()

	if resp.StatusCode != http.StatusOK {
		var errorData []byte
		errorData, err = io.ReadAll(resp.Body)
		if err != nil {
			return GatewayWebsocketInformation{}, fmt.Errorf("could not readall data from resp.Body (%w)", err)
		}
		return GatewayWebsocketInformation{}, WebAPIError{JSONData: errorData} //todo: convert to real error type
	}

	var gwi GatewayWebsocketInformation

	decoder := json.NewDecoder(resp.Body)
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&gwi)
	if err != nil {
		return GatewayWebsocketInformation{}, fmt.Errorf("could not decode json (%w)", err)
	}

	err = resp.Body.Close()
	if err != nil {
		return GatewayWebsocketInformation{}, fmt.Errorf("could not close body (%w)", err)
	}

	return gwi, nil
}

//GetGatewayWebsocketURI returns the current Discord Gateway WSS URL, pass discordAPIGatewayURL as "" to use default  //todo: make it so test doesn't have to hit server
func GetGatewayWebsocketURI(discordAPIGatewayURL string) (url.URL, error) {
	if discordAPIGatewayURL == "" {
		discordAPIGatewayURL = "https://discord.com/api/gateway"
	}

	req, err := http.NewRequest("GET", discordAPIGatewayURL, nil)
	if err != nil {
		return url.URL{}, err
	}
	req.Header.Set("User-Agent", libinfo.BotUserAgent)

	httpClient := http.Client{}
	resp, err := httpClient.Do(req)
	if err != nil {
		return url.URL{}, err
	}

	urlJSON := struct {
		URL string `json:"url"`
	}{}

	decoder := json.NewDecoder(resp.Body)
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&urlJSON)
	if err != nil {
		err2 := resp.Body.Close()
		if err2 != nil {
			return url.URL{}, fmt.Errorf("could not close body (%v), after error (%w)", err2, err)
		}
		return url.URL{}, err
	}

	err = resp.Body.Close()
	if err != nil {
		return url.URL{}, err
	}

	uri, err := url.ParseRequestURI(urlJSON.URL)
	if err != nil {
		return url.URL{}, err
	}

	return *uri, err
}
