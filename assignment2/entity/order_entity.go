package entity

import (
	"time"
)

type Order struct {
	OrderID      uint64    `gorm:"primaryKey" json:"orderId"`
	CustomerName string    `json:"customerName"`
	OrderedAt    time.Time `json:"orderedAt"`
	Items        []Item    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"items"`
	BaseModel
}

type OrderWithPerson struct {
	OrderID      uint64    `gorm:"primaryKey" json:"orderId"`
	CustomerName string    `json:"customerName"`
	OrderedAt    time.Time `json:"orderedAt"`
	Items        []Item    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"items"`
	Person       Person
	BaseModel
}
