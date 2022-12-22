package data

import (
	"skyshi/features/activity"
	"skyshi/models"

	"gorm.io/gorm"
)

type Storage struct {
	query *gorm.DB
}

func New(db *gorm.DB) activity.DataInterface {
	return &Storage{
		query: db,
	}
}

func (storage *Storage) GetActivities() ([]activity.Activity, string, error) {
	var model []models.Activity
	tx := storage.query.Find(&model)
	if tx.Error != nil {
		return nil, "DataBase Error", tx.Error
	}

	return models.ToCoreList(model), "Success", nil
}

func (storage *Storage) GetActivity(id uint) (activity.Activity, string, error) {
	var model models.Activity

	tx := storage.query.Where("id = ?", id).First(&model)
	if tx.Error != nil {
		return activity.Activity{}, "Not Found", nil
	}

	return model.ToCore(), "Success", nil
}

func (storage *Storage) Create(core activity.Activity) (activity.Activity, string, error) {
	model := models.CoreToModel(core)
	tx := storage.query.Create(&model)
	if tx.Error != nil {
		return activity.Activity{}, "Database Error", tx.Error
	}

	return model.ToCore(), "Success", nil
}

func (storage *Storage) Delete(id uint) (string, error) {
	tx := storage.query.Where("id = ?", id).Delete(&activity.Activity{})
	if tx.Error != nil || tx.RowsAffected == 0 {
		return "Not Found", nil
	}

	return "Success", nil
}

func (storage *Storage) Update(id uint, core activity.Activity) (string, error) {
	update := models.CoreToModel(core)
	tx := storage.query.Model(&models.Activity{}).Where("id = ?", id).Updates(update)
	if tx.Error != nil || tx.RowsAffected == 0 {
		return "Not Found", nil
	}

	return "Success", nil
}
