package db

import (
	"go-ddd-quickstart/internal/pkg/dto"
	"time"

	"github.com/google/uuid"
)

type Address struct {
	dto.Address
}

type OrderItem struct {
	OrderID          uuid.UUID
	IsDispatched     bool
	Address          Address
	CreatedTimestamp time.Time
	UpdatedTimestamp time.Time
}

func (i OrderItem) ItemID() uuid.UUID {
	return i.OrderID
}

func (i OrderItem) CreatedAt() time.Time {
	return i.CreatedTimestamp
}

func (i OrderItem) UpdatedAt() time.Time {
	return i.UpdatedTimestamp
}
