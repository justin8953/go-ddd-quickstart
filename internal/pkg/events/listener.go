package events

type EventListener struct {
	handlers map[string][]EventConsumer
}

func (e *EventListener) Subscribe(handler EventConsumer, events ...IEvent) {
	for _, event := range events {
		handlers := e.handlers[event.Name()]
		handlers = append(handlers, handler)
		e.handlers[event.Name()] = handlers
	}
}

func (e *EventListener) Consume(event IEvent, callback func(event map[string]interface{}) error) {
	if event.IsAsynchronous() {
		go e.consume(event, callback) // runs code in separate Go routine
	}

	e.consume(event, callback) // synchronous call
}

func (e *EventListener) consume(event IEvent, callback func(event map[string]interface{}) error) {
	for _, handler := range e.handlers[event.Name()] {
		handler.Listen(callback)
	}
}
