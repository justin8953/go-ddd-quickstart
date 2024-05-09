package dto

import (
	"go-ddd-quickstart/internal/pkg/events"
	"go-ddd-quickstart/internal/pkg/events/order"

	"github.com/google/uuid"
)

type Address struct {
	Address1 string
	Address2 string
	City     string
	State    string
	ZipCode  string
	Country  string
}

type Order struct {
	Id              uuid.UUID
	UserID          uuid.UUID
	IsDispatched    bool
	DeliveryAddress Address
}

func (o Order) ID() uuid.UUID {
	return o.Id
}

func (o *Order) ChangeAddress(address Address) events.IEvent {
	if o.IsDispatched {
		return order.OrderDeliveryAddressChangeFailed{
			OrderId: o.ID(),
		}
	}
	return order.OrderDeliveryAddressChanged{
		OrderId: o.ID(),
	}
}
