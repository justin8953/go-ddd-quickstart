package order

import (
	"os"

	"github.com/google/uuid"
)

type OrderDelivered struct {
	OrderId uuid.UUID
}

func (e OrderDelivered) IsAsynchronous() bool {
	return os.Getenv("ASYNC") == "true"
}

func (e OrderDelivered) EventId() string {
	return e.OrderID().String()
}

func (e OrderDelivered) Name() string {
	return "event.order.delivery.success"
}

func (e OrderDelivered) OrderID() uuid.UUID {
	return e.OrderId
}
