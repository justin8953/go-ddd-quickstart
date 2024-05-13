package order

import (
	orderDto "go-ddd-quickstart/internal/pkg/dto"
	"go-ddd-quickstart/internal/pkg/events"
	orderEvents "go-ddd-quickstart/internal/pkg/events/order"
	dbRecord "go-ddd-quickstart/internal/pkg/order/db"

	"github.com/google/uuid"
)

type OrderService struct {
	repository OrderRepository
	publisher  events.EventPublisher
}

func (s *OrderService) Create(order *dbRecord.OrderItem) (*dbRecord.OrderItem, error) {
	result, err := s.repository.Create(*order)
	if err != nil {
		return nil, err
	}
	//
	// update Adrress in DB
	//
	s.publisher.Notify(orderEvents.OrderCreated{
		OrderId: result.ItemID(),
	})

	return result, err
}

func (s *OrderService) ChangeAddress(userID uuid.UUID, record *dbRecord.OrderItem, address dbRecord.Address) {
	updateAddr := orderDto.Address{
		Address1: address.Address1,
		Address2: address.Address2,
		City:     address.City,
		State:    address.State,
		Country:  address.Country,
		ZipCode:  address.ZipCode,
	}
	order := orderDto.Order{
		Id:              record.ItemID(),
		UserID:          userID,
		IsDispatched:    record.IsDispatched,
		DeliveryAddress: updateAddr,
	}
	event := order.ChangeAddress(updateAddr)
	s.publisher.Notify(event) // publishing of events only inside stateless objects
}
