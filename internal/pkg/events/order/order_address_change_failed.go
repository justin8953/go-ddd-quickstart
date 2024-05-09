package order

import (
	"os"

	"github.com/google/uuid"
)

type OrderDeliveryAddressChangeFailed struct {
	OrderId uuid.UUID
}

func (e OrderDeliveryAddressChangeFailed) IsAsynchronous() bool {
	return os.Getenv("ASYNC") == "true"
}

func (e OrderDeliveryAddressChangeFailed) Name() string {
	return "event.order.address.change.failed"
}

func (e OrderDeliveryAddressChangeFailed) OrderID() uuid.UUID {
	return e.OrderId
}
