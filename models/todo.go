package models

import (
	"skyshi/features/todo"

	"gorm.io/gorm"
)

type Todo struct {
	gorm.Model
	Activity_group_id uint
	Title             string
	Is_Active         bool   `gorm:"default:true"`
	Priority          string `gorm:"default:very-high"`
}

func (a *Todo) ToCore() todo.Todo {
	return todo.Todo{
		ID:                a.ID,
		Activity_group_id: a.Activity_group_id,
		Title:             a.Title,
		Is_Active:         a.Is_Active,
		Priority:          a.Priority,
		Created_at:        a.CreatedAt,
		Updated_at:        a.UpdatedAt,
		Deleted_at:        a.DeletedAt.Time,
	}
}

func ToCoreListTodo(data []Todo) []todo.Todo {
	var result []todo.Todo
	for _, v := range data {
		result = append(result, v.ToCore())
	}

	return result
}

func CoreToModelTodo(data todo.Todo) Todo {
	return Todo{
		Activity_group_id: data.Activity_group_id,
		Title:             data.Title,
	}
}
