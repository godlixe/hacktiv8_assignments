package entity

import (
	"jwt-hacktiv8/helpers"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type User struct {
	GormModel
	FullName string    `gorm:"not null" json:"full_name" form:"full_name" validation:"required"`
	Email    string    `gorm:"not null;uniqueIndex" json:"email" form:"email" validation:"required"`
	Password string    `gorm:"not null" json:"password" binding:"validation,min=6"`
	Products []Product `gorm:"constraint:OnUpdate:CASCADE, OnDelete:SET NULL, OnDelete:SET NULL;" json:"products"`
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	_, err := govalidator.ValidateStruct(u)
	if err != nil {
		return err
	}

	u.Password, err = helpers.HashAndSalt(u.Password)
	if err != nil {
		return err
	}
	return nil
}
