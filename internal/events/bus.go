package events

import (
	"log"
	"sync"
)

type Event interface {
	Name() string
}

type EventHandler func(Event)

type EventBus struct {
	handlers map[string][]EventHandler
	mu       sync.RWMutex
	wg       sync.WaitGroup
}

func NewEventBus() *EventBus {
	return &EventBus{
		handlers: make(map[string][]EventHandler),
	}
}

func (eb *EventBus) Register(eventName string, handler EventHandler) {
	eb.mu.Lock()
	defer eb.mu.Unlock()

	eb.handlers[eventName] = append(eb.handlers[eventName], handler)
}

func (eb *EventBus) Publish(event Event) {
	eb.mu.RLock()
	defer eb.mu.RUnlock()

	if handlers, ok := eb.handlers[event.Name()]; ok {
		for _, handler := range handlers {
			eb.wg.Add(1)
			go func(h EventHandler) {
				defer eb.wg.Done()
				defer func() {
					if r := recover(); r != nil {
						log.Printf("Recovered from panic in handler for %s: %v", event.Name(), r)
					}
				}()
				h(event)
			}(handler)
		}
	}
}

func (eb *EventBus) Wait() {
	eb.wg.Wait()
}
