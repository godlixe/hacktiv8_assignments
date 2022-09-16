package dto

import "order_service/entity"

type GidhanAPIResponseDTO struct {
	Status struct {
		Code        int    `json:"code"`
		Description string `json:"description"`
	} `json:"status"`
	Result []entity.Person `json:"result"`
}
