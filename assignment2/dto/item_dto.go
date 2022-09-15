package dto

type ItemInsertDTO struct {
	ItemCode    string `json:"itemCode"`
	OrderID     uint64 `json:"orderId"`
	Description string `json:"description"`
	Quantity    uint64 `json:"quantity"`
}

type ItemUpdateDTO struct {
	ID          uint64 `json:"itemId"`
	ItemCode    string `json:"itemCode"`
	OrderID     uint64 `json:"orderId"`
	Description string `json:"description"`
	Quantity    uint64 `json:"quantity"`
}
