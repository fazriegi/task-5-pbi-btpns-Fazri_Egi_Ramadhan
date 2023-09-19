package queries

import (
	"task-5-pbi-btpns-Fazri_Egi_Ramadhan/database"
	"task-5-pbi-btpns-Fazri_Egi_Ramadhan/models"
)

type PhotoQuery struct {}
func (p *PhotoQuery) Get(userId uint) (uint, error) {
	var id uint
	if err := database.DB.Table("photos").Select("id").Where("user_id = ?", userId).Scan(&id).Error; err != nil {
		return 0, err
	}

	return id, nil
}

func (p *PhotoQuery) Delete(photoId uint) error {
	if err := database.DB.Delete(&models.Photo{}, photoId).Error; err != nil {
		return err
	}

	return nil
}