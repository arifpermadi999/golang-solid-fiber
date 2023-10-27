package models

import "gorm.io/gorm"

type Unit struct {
	Id   int    `json:"id"`
	Name string `gorm:"type:varchar(255)" json:"name" valid:"required"`

	gorm.Model
}
