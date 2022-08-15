package dizzy

import (
	"sync"

	"github.com/SomethingBot/dizzy/primitives"
)

type eDist interface {
	FireEvent(event primitives.GatewayEvent)
	RegisterHandler(eventType primitives.GatewayEventType, handler EventHandler) DeregisterFunction
	RegisterHandlerOnce(eventType primitives.GatewayEventType, handler EventHandler)
	WaitTilDone()
}

// EventHandler function that is called on a event
type EventHandler func(event primitives.GatewayEvent)

type handlerEntry struct { //todo: maybe try letting this be locked separately from map containing it
	handlers   []EventHandler
	nilIndexes []int
}

// EventDistributor for events
type EventDistributor struct {
	sync.RWMutex
	runningHandlers sync.WaitGroup
	handlerEntries  map[primitives.GatewayEventType]*handlerEntry
}

// NewEventDistributor for events
func NewEventDistributor() *EventDistributor {
	return &EventDistributor{handlerEntries: make(map[primitives.GatewayEventType]*handlerEntry)}
}

// FireEvent to all listening EventHandler(s)
func (e *EventDistributor) FireEvent(event primitives.GatewayEvent) {
	e.runningHandlers.Add(1)
	go func() {
		defer e.runningHandlers.Done()
		e.RWMutex.RLock()
		handlerEntry, ok := e.handlerEntries[event.Type()]
		if !ok {
			return
		}
		for _, handler := range handlerEntry.handlers {
			if handler != nil {
				e.runningHandlers.Add(1)
				go func(eventHandler EventHandler) {
					eventHandler(event)
					e.runningHandlers.Done()
				}(handler)
			}
		}
		e.RWMutex.RUnlock()
	}()
}

// DeregisterFunction is a function that is called to remove an EventHandler from the EventDistributor
type DeregisterFunction func()

type deregisterEntry struct {
	sync.Mutex
	fired            bool
	eventType        primitives.GatewayEventType
	index            int
	eventDistributor *EventDistributor
}

func (d *deregisterEntry) deregisterHandler() {
	d.Lock()
	defer d.Unlock()
	if d.fired {
		return
	}
	d.fired = true

	d.eventDistributor.RWMutex.Lock()
	defer d.eventDistributor.RWMutex.Unlock()
	entry := d.eventDistributor.handlerEntries[d.eventType]
	entry.handlers[d.index] = nil
	entry.nilIndexes = append(entry.nilIndexes, d.index)
	if len(entry.nilIndexes) == len(entry.handlers) { //todo: this allows us to reset memory usage of slices if we eventually drain all of the handlers, however this is only when *every* handler is removed
		entry.handlers = []EventHandler{}
		entry.nilIndexes = []int{}
	}
}

// RegisterHandler to be called when an Event is fired
func (e *EventDistributor) RegisterHandler(eventType primitives.GatewayEventType, handler EventHandler) DeregisterFunction {
	deregEntry := deregisterEntry{}
	e.RWMutex.Lock()
	entry, ok := e.handlerEntries[eventType]
	if !ok {
		e.handlerEntries[eventType] = &handlerEntry{
			handlers: []EventHandler{handler},
		}
		e.RWMutex.Unlock()
		deregEntry = deregisterEntry{
			eventType:        eventType,
			index:            0,
			eventDistributor: e,
		}
		return deregEntry.deregisterHandler
	}
	if len(entry.nilIndexes) > 0 {
		index := entry.nilIndexes[0]
		entry.handlers[index] = handler
		var slice []int
		slice = append(slice, entry.nilIndexes[1:]...)
		entry.nilIndexes = slice
		e.RWMutex.Unlock()
		deregEntry = deregisterEntry{
			eventType:        eventType,
			index:            index,
			eventDistributor: e,
		}
		return deregEntry.deregisterHandler
	}
	entry.handlers = append(entry.handlers, handler)
	index := len(entry.handlers) - 1

	e.RWMutex.Unlock()
	deregEntry = deregisterEntry{
		eventType:        eventType,
		index:            index,
		eventDistributor: e,
	}
	return deregEntry.deregisterHandler
}

// WaitTilDone with all Events
func (e *EventDistributor) WaitTilDone() {
	e.runningHandlers.Wait()
}

type singleFireHandler struct {
	sync.Mutex
	fired              bool
	handler            EventHandler
	deregisterFunction DeregisterFunction
}

func (sfh *singleFireHandler) Fire(event primitives.GatewayEvent) {
	sfh.Lock()
	if sfh.fired {
		sfh.Unlock()
		return
	}
	sfh.handler(event)
	sfh.fired = true
	sfh.deregisterFunction()
	sfh.Unlock()
}

// RegisterHandlerOnce lets an EventHandler be called and then removed after being called exactly once
func (e *EventDistributor) RegisterHandlerOnce(eventType primitives.GatewayEventType, handler EventHandler) {
	sfh := singleFireHandler{handler: handler}
	sfh.deregisterFunction = e.RegisterHandler(eventType, sfh.Fire)
}
