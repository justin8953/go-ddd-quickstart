package notification

import (
	"os"

	"github.com/google/uuid"
)

type EmailSent struct {
	EmailId uuid.UUID
}

func (e EmailSent) IsAsynchronous() bool {
	return os.Getenv("ASYNC") == "true"
}
func (e EmailSent) Name() string {
	return "event.email.sent"
}

func (e EmailSent) EmailID() uuid.UUID {
	return e.EmailId
}
