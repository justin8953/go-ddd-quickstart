package notification

import (
	"go-ddd-quickstart/internal/pkg/events"

	"github.com/google/uuid"
)

type EmailEvent interface {
	events.IEvent
	EmailID() uuid.UUID
}
