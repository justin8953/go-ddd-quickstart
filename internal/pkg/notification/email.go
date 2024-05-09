package notification

import (
	"go-ddd-quickstart/internal/pkg/events"
	emailEvent "go-ddd-quickstart/internal/pkg/events/notification"
)

type EmailHandler struct {
}

func (e *EmailHandler) handleEmailSent(event emailEvent.EmailSent) {
	// Do something with the event
}

func (e *EmailHandler) Notify(event events.IEvent) {
	switch actualEvent := event.(type) {
	case emailEvent.EmailSent:
		e.handleEmailSent(actualEvent)
	default:
		return
	}
}
