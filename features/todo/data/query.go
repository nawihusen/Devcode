package data

import (
	"skyshi/features/todo"
	"skyshi/models"

	"gorm.io/gorm"
)

type Storage struct {
	query *gorm.DB
}

func New(db *gorm.DB) todo.DataInterface {
	return &Storage{
		query: db,
	}
}

func (storage *Storage) GetAllTodo() ([]todo.Todo, string, error) {
	var model []models.Todo
	tx := storage.query.Find(&model)
	if tx.Error != nil {
		return nil, "DataBase Error", tx.Error
	}

	return models.ToCoreListTodo(model), "Success", nil
}

func (storage *Storage) GetATodo(id uint) (todo.Todo, string, error) {
	var model models.Todo

	tx := storage.query.Where("id = ?", id).First(&model)
	if tx.Error != nil {
		return todo.Todo{}, "DataBase Error", tx.Error
	} else if tx.RowsAffected == 0 {
		return todo.Todo{}, "Not Found", nil
	}

	return model.ToCore(), "Success", nil
}

func (storage *Storage) Create(core todo.Todo) (todo.Todo, string, error) {
	model := models.CoreToModelTodo(core)
	tx := storage.query.Create(&model)
	if tx.Error != nil {
		return todo.Todo{}, "Database Error", tx.Error
	}
	return model.ToCore(), "Success", nil
}

func (storage *Storage) Delete(id uint) (string, error) {
	tx := storage.query.Where("id = ?", id).Delete(todo.Todo{})
	if tx.Error != nil {
		return "DataBase Error", tx.Error
	} else if tx.RowsAffected == 0 {
		return "Not Found", nil
	}

	return "Success", nil
}

func (storage *Storage) Update(id uint, core todo.Todo) (string, error) {
	update := models.CoreToModelTodo(core)
	tx := storage.query.Model(&models.Todo{}).Where("id = ?", id).Updates(update)
	if tx.Error != nil {
		return "DataBase Error", tx.Error
	} else if tx.RowsAffected == 0 {
		return "Not Found", nil
	}

	return "Success", nil
}
