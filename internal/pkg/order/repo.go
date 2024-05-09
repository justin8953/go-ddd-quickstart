package order

import (
	"go-ddd-quickstart/internal/pkg/db"
	dbRecord "go-ddd-quickstart/internal/pkg/order/db"

	"github.com/google/uuid"
)

type OrderRepository struct {
	db.DbRepo
}

func (r *OrderRepository) Create(item *dbRecord.OrderItem) (*dbRecord.OrderItem, error) {
	return &dbRecord.OrderItem{
		OrderID:          item.OrderID,
		CreatedTimestamp: item.CreatedTimestamp,
		UpdatedTimestamp: item.UpdatedTimestamp,
	}, nil
}

func (r *OrderRepository) Update(id uuid.UUID, item *dbRecord.OrderItem) (*dbRecord.OrderItem, error) {
	return &dbRecord.OrderItem{
		OrderID:          item.OrderID,
		CreatedTimestamp: item.CreatedTimestamp,
		UpdatedTimestamp: item.UpdatedTimestamp,
	}, nil
}

func (r *OrderRepository) Delete(id uuid.UUID) error {
	return nil
}

func (r *OrderRepository) List(filter map[string]interface{}) ([]dbRecord.OrderItem, error) {
	return nil, nil
}

func (r *OrderRepository) Retrieve(id uuid.UUID) (*dbRecord.OrderItem, error) {
	return nil, nil
}
