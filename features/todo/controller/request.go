package controller

import "skyshi/features/todo"

type Request struct {
	Activity_group_id uint   `json:"activity_group_id"`
	Title             string `json:"title"`
}

type UpdateRequest struct {
	Activity_group_id uint   `json:"activity_group_id"`
	Title             string `json:"title"`
	Is_active         bool   `json:"is_active"`
	Priority          string `json:"priority"`
}

func (r *Request) ToCore() todo.Todo {
	return todo.Todo{
		Activity_group_id: r.Activity_group_id,
		Title:             r.Title,
	}
}

func (r *UpdateRequest) ToCore() todo.Todo {
	return todo.Todo{
		Activity_group_id: r.Activity_group_id,
		Title:             r.Title,
		Is_Active:         r.Is_active,
		Priority:          r.Priority,
	}
}
