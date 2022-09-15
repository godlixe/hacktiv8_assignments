package entity

import "time"

// Base model that includes uint64 ID and created, updated, deleted timestamps
type BaseModel struct {
	ID        uint64
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}
