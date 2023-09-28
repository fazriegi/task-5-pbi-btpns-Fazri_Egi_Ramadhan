package models

import "gorm.io/gorm"

type Photo struct {
	gorm.Model
	Title    string `form:"title" gorm:"type:varchar(100)"`
	Caption  string `form:"caption"`
	PhotoURL string `form:"photo_url"`
	UserID   uint   `form:"user_id"`
	User     User   `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
}
