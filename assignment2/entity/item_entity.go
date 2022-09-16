package entity

type Item struct {
	ItemID      uint64 `gorm:"primaryKey" json:"lineItemId"`
	ItemCode    string `json:"itemCode"`
	OrderID     uint64 `json:"orderId"`
	Order       *Order `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"order"`
	Description string `json:"description"`
	Quantity    uint64 `json:"quantity"`
	BaseModel
}
