package order

import (
	"os"

	"github.com/google/uuid"
)

type OrderDeliveryAddressChanged struct {
	OrderId uuid.UUID
}

func (e OrderDeliveryAddressChanged) IsAsynchronous() bool {
	return os.Getenv("ASYNC") == "true"
}

func (e OrderDeliveryAddressChanged) EventId() string {
	return e.OrderID().String()
}

func (e OrderDeliveryAddressChanged) Name() string {
	return "event.order.address.change.success"
}

func (e OrderDeliveryAddressChanged) OrderID() uuid.UUID {
	return e.OrderId
}
