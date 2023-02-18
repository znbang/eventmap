package bookservice

import (
	"sync"
)

var EventBus = newEventBus()

type Event struct {
	ID      string `json:"id"`
	Message string `json:"message"`
	Status  int    `json:"status"`
}

type eventBus struct {
	subscribers map[chan Event]bool
	mutex       *sync.RWMutex
}

func newEventBus() *eventBus {
	return &eventBus{
		subscribers: make(map[chan Event]bool),
		mutex:       &sync.RWMutex{},
	}
}

func (e *eventBus) Register() chan Event {
	e.mutex.Lock()
	defer e.mutex.Unlock()

	c := make(chan Event, 10)
	e.subscribers[c] = true
	return c
}

func (e *eventBus) Unregister(c chan Event) {
	e.mutex.Lock()
	defer e.mutex.Unlock()

	close(c)
	delete(e.subscribers, c)
}

func (e *eventBus) Post(event Event) error {
	e.mutex.RLock()
	defer e.mutex.RUnlock()

	for c := range e.subscribers {
		select {
		case c <- event:
		default:
		}
	}

	return nil
}

func (e *eventBus) Close() {
	e.mutex.Lock()
	defer e.mutex.Unlock()

	for c := range e.subscribers {
		close(c)
		delete(e.subscribers, c)
	}
}
