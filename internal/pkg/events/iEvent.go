package events

type IEvent interface {
	Name() string
	IsAsynchronous() bool
}
