package migrations

import (
	"skyshi/models"

	"gorm.io/gorm"
)

func InitMigrate(db *gorm.DB) {
	db.AutoMigrate(models.Activity{})
	db.AutoMigrate(models.Todo{})
}
