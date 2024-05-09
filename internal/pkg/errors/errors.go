package errors

import "fmt"

type EventError struct {
	EventName string
	Err       error
}

func NewEventError(eventName string, err error) EventError {
	return EventError{
		EventName: eventName,
		Err:       err,
	}
}

func (e EventError) Name() string {
	return e.EventName
}

func (e EventError) Error() string {
	return fmt.Sprintf("event name %s: err %v", e.EventName, e.Err)
}
