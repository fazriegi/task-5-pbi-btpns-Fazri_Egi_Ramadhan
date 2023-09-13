package queries

import (
	"task-5-pbi-btpns-Fazri_Egi_Ramadhan/database"
	"task-5-pbi-btpns-Fazri_Egi_Ramadhan/models"
)

func Save(user *models.User) error {
	return database.DB.Save(user).Error
}
