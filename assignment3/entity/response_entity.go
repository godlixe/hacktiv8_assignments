package entity

import "time"

type Response struct {
	Data      Data       `json:"data"`
	Status    DataStatus `json:"status"`
	UpdatedAt time.Time  `json:"updatedAt"`
}
