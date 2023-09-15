package queries

import (
	"task-5-pbi-btpns-Fazri_Egi_Ramadhan/database"
	"task-5-pbi-btpns-Fazri_Egi_Ramadhan/models"

	"gorm.io/gorm"
)

func Save(user *models.User) error {
	return database.DB.Save(user).Error
}

func GetUser(email string) (models.User, error) {
	var err error
	var user models.User

	if err = database.DB.Table("users").Where("email = ?", email).First(&user).Error; err != nil && err != gorm.ErrRecordNotFound {
		return user, err
	}

	if err == gorm.ErrRecordNotFound {
		return user, nil
	}

	return user, nil
}
