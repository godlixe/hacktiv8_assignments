package entity

type Item struct {
	BaseModel
	ItemCode    string
	OrderID     uint64
	Order       *Order `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Description string
	Quantity    uint64
}
