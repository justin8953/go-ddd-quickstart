package events

type EventHandler interface {
	Topic() string
	Notify(event IEvent)
}

type EventConsumer interface {
	Topic() string
	Listen(callback func(event map[string]interface{}) error)
}
