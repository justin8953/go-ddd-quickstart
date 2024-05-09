package events

type IEvent interface {
	Name() string
	EventId() string
	IsAsynchronous() bool
}
