package todo

import "time"

type Todo struct {
	ID                uint
	Activity_group_id uint
	Title             string
	Is_Active         bool
	Priority          string
	Created_at        time.Time
	Updated_at        time.Time
	Deleted_at        time.Time
}

type ServiceInterface interface {
	GetAllTodo() ([]Todo, string, error)
	GetATodo(id uint) (Todo, string, error)
	Create(Todo) (Todo, string, error)
	Delete(id uint) (string, error)
	Update(id uint, data Todo) (Todo, string, error)
}

type DataInterface interface {
	GetAllTodo() ([]Todo, string, error)
	GetATodo(id uint) (Todo, string, error)
	Create(Todo) (Todo, string, error)
	Delete(id uint) (string, error)
	Update(id uint, data Todo) (string, error)
}
