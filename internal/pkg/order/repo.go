package order

import (
	"go-ddd-quickstart/internal/pkg/mongo_connector"
	dbRecord "go-ddd-quickstart/internal/pkg/order/db"
	"time"

	"github.com/google/uuid"
)

type OrderRepository struct {
	Repo *mongo_connector.MongoService
}

func NewOrderRepository() *OrderRepository {
	repo := mongo_connector.NewMongoService([]string{"order"})
	return &OrderRepository{
		Repo: repo,
	}
}

func (r *OrderRepository) Create(item *dbRecord.OrderItem) (*dbRecord.OrderItem, error) {
	item.CreatedTimestamp = time.Now()
	item.UpdatedTimestamp = time.Now()
	id, err := r.Repo.GetMongoRepo("order").Create(item)
	if err != nil {
		return nil, err
	}
	newItem, err := r.Repo.GetMongoRepo("order").Retrieve(uuid.MustParse(id))
	if err != nil {
		return nil, err
	}
	return newItem.(*dbRecord.OrderItem), nil
}

func (r *OrderRepository) Update(id uuid.UUID, item *dbRecord.OrderItem) (*dbRecord.OrderItem, error) {
	item.UpdatedTimestamp = time.Now()
	err := r.Repo.GetMongoRepo("order").Update(id, item)
	if err != nil {
		return nil, err
	}
	updateItem, err := r.Repo.GetMongoRepo("order").Retrieve(id)
	if err != nil {
		return nil, err
	}
	return updateItem.(*dbRecord.OrderItem), nil
}

func (r *OrderRepository) Delete(id uuid.UUID) error {
	err := r.Repo.GetMongoRepo("order").Delete(id)
	if err != nil {
		return err
	}
	return nil
}

func (r *OrderRepository) List(filter map[string]interface{}) ([]dbRecord.OrderItem, error) {
	cursor, err := r.Repo.GetMongoRepo("order").List(filter)
	if err != nil {
		return nil, err
	}
	var items []dbRecord.OrderItem
	for _, item := range cursor {
		items = append(items, *item.(*dbRecord.OrderItem))
	}
	return items, nil
}

func (r *OrderRepository) Retrieve(id uuid.UUID) (*dbRecord.OrderItem, error) {
	item, err := r.Repo.GetMongoRepo("order").Retrieve(id)
	if err != nil {
		return nil, err
	}
	return item.(*dbRecord.OrderItem), nil
}
