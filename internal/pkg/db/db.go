package db

type DbRepo interface {
	Create(item IItem) (string, error)
	Update(id string, item IItem) error
	Delete(id string) error
	List(filter map[string]interface{}) ([]IItem, error)
	Retrieve(id string) (IItem, error)
}
