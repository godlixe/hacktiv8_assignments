package dto

import (
	"time"
)

type OrderCreateDTO struct {
	CustomerName string          `json:"customerName" binding:"required"`
	Items        []ItemInsertDTO `json:"items" binding:"required"`
	OrderedAt    time.Time       `json:"orderedAt" binding:"required"`
}

type OrderUpdateDTO struct {
	OrderID      uint64          `json:"orderId" binding:"required"`
	CustomerName string          `json:"customerName" binding:"required"`
	Items        []ItemUpdateDTO `json:"items" binding:"required"`
	OrderedAt    time.Time       `json:"orderedAt" binding:"required"`
}
