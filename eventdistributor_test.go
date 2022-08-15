package discord

import (
	"sync"
	"testing"
	"time"

	"github.com/SomethingBot/dizzy/primitives"
)

type testHandlerUser struct {
	sync.Mutex
	count     int
	firedChan chan bool
}

func (th *testHandlerUser) Handle(event primitives.GatewayEvent) {
	th.Lock()
	th.count++
	th.Unlock()
	th.firedChan <- true
}

func TestGatewayEventHandler_FireEvent(t *testing.T) {
	t.Parallel()

	eventDistributor := NewEventDistributor()

	thu := &testHandlerUser{firedChan: make(chan bool)}
	eventDistributor.RegisterHandler(primitives.GatewayEventTypeHello, thu.Handle)

	handlerUserCount := 100

	for i := 0; i < handlerUserCount; i++ {
		eventDistributor.FireEvent(primitives.GatewayEventHello{})
	}

	timer := time.NewTimer(time.Second)
	for i := 0; i < handlerUserCount; i++ {
		select {
		case fired := <-thu.firedChan:
			if !fired {
				t.Fatalf("how did we get here")
			}
			timer.Reset(time.Second)
		case <-timer.C:
			t.Fatalf("event fire timed out")
		}
	}

	if thu.count != handlerUserCount {
		t.Fatalf("count is: (%v), wanted count (%v)\n", thu.count, handlerUserCount)
	}

	firedChan := make(chan bool)
	firedChan2 := make(chan bool)
	eventDistributor.RegisterHandler(primitives.GatewayEventTypeHello, func(hello primitives.GatewayEvent) {
		firedChan <- true
	})
	eventDistributor.RegisterHandler(primitives.GatewayEventTypeHello, func(hello primitives.GatewayEvent) {
		firedChan2 <- true
	})

	count := 10

	for i := 0; i < count; i++ {
		go eventDistributor.FireEvent(primitives.GatewayEventHello{})
	}

	timer = time.NewTimer(time.Second)
	for i := 0; i < count*2; i++ {
		select {
		case fired := <-firedChan:
			if !fired {
				t.Fatalf("how did we get here")
			}
			timer.Reset(time.Second)
		case fired := <-firedChan2:
			if !fired {
				t.Fatalf("how did we get here")
			}
			timer.Reset(time.Second)
		case <-timer.C:
			t.Fatalf("event fire timed out")
		}
	}
}
