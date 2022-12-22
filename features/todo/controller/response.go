package controller

import (
	"skyshi/features/todo"
	"strconv"
	"time"
)

type Response struct {
	ID                uint      `json:"id"`
	Activity_group_id string    `json:"activity_group_id"`
	Title             string    `json:"title"`
	Is_Active         string    `json:"is_active"`
	Priority          string    `json:"priority"`
	Created_at        time.Time `json:"created_at"`
	Updated_at        time.Time `json:"updated_at"`
	Deleted_at        []int     `json:"deleted_at"`
}

type CreateResponse struct {
	Created_at        time.Time `json:"created_at"`
	Updated_at        time.Time `json:"updated_at"`
	ID                uint      `json:"id"`
	Title             string    `json:"title"`
	Activity_group_id uint      `json:"activity_group_id"`
	Is_Active         bool      `json:"is_active"`
	Priority          string    `json:"priority"`
}

func CoreToResponse(data todo.Todo) Response {
	idres := strconv.Itoa(int(data.Activity_group_id))
	return Response{
		ID:                data.ID,
		Activity_group_id: idres,
		Title:             data.Title,
		Is_Active:         "1",
		Priority:          data.Priority,
		Created_at:        data.Created_at,
		Updated_at:        data.Updated_at,
		Deleted_at:        nil,
	}
}

func CoreToResponseList(data []todo.Todo) []Response {
	var res []Response
	for _, v := range data {
		res = append(res, CoreToResponse(v))
	}

	return res
}

func CoreToCreaterResponse(data todo.Todo) CreateResponse {
	return CreateResponse{
		Created_at:        data.Created_at,
		Updated_at:        data.Updated_at,
		ID:                data.ID,
		Title:             data.Title,
		Activity_group_id: data.Activity_group_id,
		Is_Active:         data.Is_Active,
		Priority:          data.Priority,
	}
}
