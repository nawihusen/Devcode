package controller

import (
	"skyshi/features/activity"
	"time"
)

type Response struct {
	ID         uint      `json:"id"`
	Email      string    `json:"email"`
	Title      string    `json:"title"`
	Created_at time.Time `json:"created_at"`
	Updated_at time.Time `json:"updated_at"`
	Deleted_at time.Time `json:"deleted_at"`
}

type CreateResponse struct {
	Created_at time.Time `json:"created_at"`
	Updated_at time.Time `json:"updated_at"`
	ID         uint      `json:"id"`
	Title      string    `json:"title"`
	Email      string    `json:"email"`
}

func CoreToResponse(data activity.Activity) Response {
	return Response{
		ID:         data.ID,
		Email:      data.Email,
		Title:      data.Title,
		Created_at: data.Created_at,
		Updated_at: data.Updated_at,
		Deleted_at: data.Deleted_at,
	}
}

func CoreToResponseList(data []activity.Activity) []Response {
	var res []Response
	for _, v := range data {
		res = append(res, CoreToResponse(v))
	}

	return res
}

func CoreToCreaterResponse(data activity.Activity) CreateResponse {
	return CreateResponse{
		Created_at: data.Created_at,
		Updated_at: data.Updated_at,
		ID:         data.ID,
		Title:      data.Title,
		Email:      data.Email,
	}
}
