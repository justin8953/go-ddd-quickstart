package db

import "github.com/google/uuid"

type DbRepo interface {
	Create(item *IItem) (*IItem, error)
	Update(id uuid.UUID, item *IItem) (*IItem, error)
	Delete(id uuid.UUID) error
	List(filter map[string]interface{}) ([]IItem, error)
	Retrieve(id uuid.UUID) (*IItem, error)
}
