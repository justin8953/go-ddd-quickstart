package events

type EventHandler interface {
	Notify(event IEvent)
}
