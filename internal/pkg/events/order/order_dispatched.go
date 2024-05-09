package order

import (
	"os"

	"github.com/google/uuid"
)

type OrderDispatched struct {
	OrderId uuid.UUID
}

func (e OrderDispatched) IsAsynchronous() bool {
	return os.Getenv("ASYNC") == "true"
}
func (e OrderDispatched) EventId() string {
	return e.OrderID().String()
}

func (e OrderDispatched) Name() string {
	return "event.order.dispatched"
}

func (e OrderDispatched) OrderID() uuid.UUID {
	return e.OrderId
}
