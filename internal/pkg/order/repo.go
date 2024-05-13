package order

import (
	"go-ddd-quickstart/internal/pkg/db"
	"go-ddd-quickstart/internal/pkg/mongo_connector"
	dbRecord "go-ddd-quickstart/internal/pkg/order/db"
	"time"
)

type OrderRepository struct {
	Repo db.DbRepo
}

func NewOrderRepository() *OrderRepository {
	service := mongo_connector.NewMongoService([]string{"order"})
	repo := service.GetMongoRepo("order")
	return &OrderRepository{
		Repo: repo,
	}
}

func (r *OrderRepository) Create(item dbRecord.OrderItem) (*dbRecord.OrderItem, error) {
	item.CreatedTimestamp = time.Now()
	item.UpdatedTimestamp = time.Now()
	id, err := r.Repo.Create(item)
	if err != nil {
		return nil, err
	}
	newItem, err := r.Repo.Retrieve(id)
	if err != nil {
		return nil, err
	}
	newOrderItem := newItem.(dbRecord.OrderItem)
	return &newOrderItem, nil
}

func (r *OrderRepository) Update(id string, item *dbRecord.OrderItem) (*dbRecord.OrderItem, error) {
	item.UpdatedTimestamp = time.Now()
	err := r.Repo.Update(id, item)
	if err != nil {
		return nil, err
	}
	updateItem, err := r.Repo.Retrieve(id)
	if err != nil {
		return nil, err
	}
	return updateItem.(*dbRecord.OrderItem), nil
}

func (r *OrderRepository) Delete(id string) error {
	err := r.Repo.Delete(id)
	if err != nil {
		return err
	}
	return nil
}

func (r *OrderRepository) List(filter map[string]interface{}) ([]dbRecord.OrderItem, error) {
	cursor, err := r.Repo.List(filter)
	if err != nil {
		return nil, err
	}
	var items []dbRecord.OrderItem
	for _, item := range cursor {
		items = append(items, *item.(*dbRecord.OrderItem))
	}
	return items, nil
}

func (r *OrderRepository) Retrieve(id string) (*dbRecord.OrderItem, error) {
	item, err := r.Repo.Retrieve(id)
	if err != nil {
		return nil, err
	}
	orderItem := item.(dbRecord.OrderItem)
	return &orderItem, nil
}
