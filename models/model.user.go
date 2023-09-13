package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `form:"username" gorm:"type:varchar(255);not null"`
	Email    string `form:"email" gorm:"type:varchar(100);unique;not null"`
	Password string `form:"password" gorm:"not null"`
}
