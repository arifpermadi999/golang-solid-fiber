package models

import "gorm.io/gorm"

type Category struct {
	Id   int    `json:"id" form:"id"`
	Name string `gorm:"type:varchar(255)" json:"name" form:"name" valid:"required"`

	gorm.Model
}
