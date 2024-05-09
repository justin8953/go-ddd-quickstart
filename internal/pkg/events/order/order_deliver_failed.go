package order

import (
	"os"

	"github.com/google/uuid"
)

type OrderDeliveryFailed struct {
	OrderId uuid.UUID
}

func (e OrderDeliveryFailed) IsAsynchronous() bool {
	return os.Getenv("ASYNC") == "true"
}

func (e OrderDeliveryFailed) Name() string {
	return "event.order.delivery.failed"
}

func (e OrderDeliveryFailed) OrderID() uuid.UUID {
	return e.OrderId
}
