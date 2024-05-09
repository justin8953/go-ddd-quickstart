package notification

import (
	"os"

	"github.com/google/uuid"
)

type EmailSendFailed struct {
	EmailId uuid.UUID
}

func (e EmailSendFailed) IsAsynchronous() bool {
	return os.Getenv("ASYNC") == "true"
}

func (e EmailSendFailed) EventId() string {
	return e.EmailID().String()
}

func (e EmailSendFailed) Name() string {
	return "event.email.sent"
}

func (e EmailSendFailed) EmailID() uuid.UUID {
	return e.EmailId
}
