package db

import (
	"time"

	"github.com/google/uuid"
)

type IItem interface {
	ItemID() uuid.UUID
	CreatedAt() time.Time
	UpdatedAt() time.Time
}
