package order

import (
	"go-ddd-quickstart/internal/pkg/events"

	"github.com/google/uuid"
)

type OrderEvent interface {
	events.IEvent
	OrderID() uuid.UUID
}
