package models

import "gorm.io/gorm"

type Photo struct {
	gorm.Model
	Title    string `json:"title" gorm:"type:varchar(100)"`
	Caption  string `json:"caption"`
	PhotoURL string `json:"photo_url"`
	UserID   uint   `json:"user_id"`
	User     User   `gorm:"foreignKey:UserID"`
}
