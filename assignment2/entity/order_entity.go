package entity

import (
	"time"
)

type Order struct {
	BaseModel
	CustomerName string
	Items        []Item `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	OrderedAt    time.Time
}
