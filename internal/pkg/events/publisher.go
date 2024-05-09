package events

type EventPublisher struct {
	handlers map[string][]EventHandler
}

func (e *EventPublisher) Subscribe(handler EventHandler, events ...IEvent) {
	for _, event := range events {
		handlers := e.handlers[event.Name()]
		handlers = append(handlers, handler)
		e.handlers[event.Name()] = handlers
	}
}

func (e *EventPublisher) Notify(event IEvent) {
	if event.IsAsynchronous() {
		go e.notify(event) // runs code in separate Go routine
	}

	e.notify(event) // synchronous call
}

func (e *EventPublisher) notify(event IEvent) {
	for _, handler := range e.handlers[event.Name()] {
		handler.Notify(event)
	}
}
