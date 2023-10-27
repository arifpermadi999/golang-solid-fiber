package models

import "gorm.io/gorm"

type User struct {
	Id       int    `json:"id"`
	Name     string `gorm:"type:varchar(255)" json:"name" valid:"required"`
	Password string `gorm:"type:varchar(255)" json:"password" valid:"required,minstringlength(8)"`
	Email    string `gorm:"uniqueIndex;type:varchar(255)" json:"email" valid:"required,email" `
	Role     string `gorm:"type:varchar(255)" json:"role" valid:"required"`

	gorm.Model
}
