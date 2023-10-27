package models

import "gorm.io/gorm"

type Product struct {
	Id    int    `json:"id"`
	Name  string `gorm:"type:varchar(255)" json:"name" valid:"required"`
	Image string `gorm:"type:text(255)" json:"image" valid:"required"`

	gorm.Model
}
