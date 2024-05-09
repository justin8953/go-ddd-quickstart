package db

import (
	"time"

	"github.com/google/uuid"
)

type IItem interface {
	ItemID() uuid.UUID
	CreateedAt() time.Time
	UpdatedAt() time.Time
}
