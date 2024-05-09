package order

import (
	"os"

	"github.com/google/uuid"
)

type OrderCreated struct {
	OrderId uuid.UUID
}

func (e OrderCreated) IsAsynchronous() bool {
	return os.Getenv("ASYNC") == "true"
}

func (e OrderCreated) EventId() string {
	return uuid.New().String()
}

func (e OrderCreated) Name() string {
	return "event.order.created"
}

func (e OrderCreated) OrderID() uuid.UUID {
	return e.OrderId
}
