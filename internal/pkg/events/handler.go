package events

type EventHandler interface {
	Topic() string
	Notify(event IEvent)
}
