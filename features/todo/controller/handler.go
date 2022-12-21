package controller

import (
	"fmt"
	"skyshi/features/todo"
	"skyshi/utils/helper"
	"strconv"

	"github.com/labstack/echo/v4"
)

type Controller struct {
	control todo.ServiceInterface
}

func New(e *echo.Echo, data todo.ServiceInterface) {
	handler := &Controller{
		control: data,
	}
	e.GET("/todo-items", handler.GetAllTodo)
	e.GET("/todo-items/:id", handler.GetTodo)
	e.POST("/todo-items", handler.Create)
	e.DELETE("/todo-items/:id", handler.Delete)
	e.PATCH("/todo-items/:id", handler.Update)
}

func (client *Controller) GetAllTodo(c echo.Context) error {
	activity_group_id := c.QueryParam("activity_group_id")

	data, msg, err := client.control.GetAllTodo()
	if err != nil {
		return c.JSON(400, helper.FailedResponseHelper(msg))
	}

	if activity_group_id != "" {
		new := []todo.Todo{}
		for _, v := range data {
			temp := strconv.Itoa(int(v.Activity_group_id))
			if temp == activity_group_id {
				new = append(new, v)
			}
		}
		return c.JSON(200, helper.SuccessDataResponseHelper(msg, CoreToResponseList(new)))
	}

	return c.JSON(200, helper.SuccessDataResponseHelper(msg, CoreToResponseList(data)))
}

func (client *Controller) GetTodo(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(400, helper.FailedResponseHelper("Parameter must be number"))
	}

	data, msg, er := client.control.GetATodo(uint(id))
	if er != nil {
		return c.JSON(400, helper.FailedResponseHelper(msg))
	} else if msg == "Not Found" {
		notfound := fmt.Sprintf("Todo with ID %d Not Found", id)
		return c.JSON(404, helper.NotFoundHelper(notfound))
	}

	return c.JSON(200, helper.SuccessDataResponseHelper(msg, CoreToResponse(data)))
}

func (client *Controller) Create(c echo.Context) error {
	var req Request
	er := c.Bind(&req)
	if er != nil {
		return c.JSON(400, helper.FailedResponseHelper("Failed Bind"))
	} else if req.Title == "" {
		return c.JSON(400, helper.BadRequestTodo("title cannot be null"))
	} else if req.Activity_group_id == 0 {
		return c.JSON(400, helper.BadRequestTodo("activity_group_id cannot be null"))
	}

	data, msg, err := client.control.Create(req.ToCore())
	if err != nil {
		return c.JSON(400, helper.FailedResponseHelper(msg))
	}

	return c.JSON(201, helper.SuccessDataResponseHelper(msg, CoreToCreaterResponse(data)))
}

func (client *Controller) Delete(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(400, helper.FailedResponseHelper("Parameter must be number"))
	}

	msg, er := client.control.Delete(uint(id))
	if er != nil {
		return c.JSON(400, helper.FailedResponseHelper(msg))
	} else if msg == "Not Found" {
		notfound := fmt.Sprintf("Todo with ID %d Not Found", id)
		return c.JSON(404, helper.NotFoundHelper(notfound))
	}

	return c.JSON(200, helper.SuccessResponseHelper(msg))
}

func (client *Controller) Update(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(400, helper.FailedResponseHelper("Parameter must be number"))
	}

	var req UpdateRequest
	er := c.Bind(&req)
	if er != nil {
		return c.JSON(400, helper.FailedResponseHelper("Failed Bind"))
	}

	data, msg, er := client.control.Update(uint(id), req.ToCore())
	if er != nil {
		return c.JSON(400, helper.FailedResponseHelper(msg))
	} else if msg == "Not Found" {
		notfound := fmt.Sprintf("Todo with ID %d Not Found", id)
		return c.JSON(404, helper.NotFoundHelper(notfound))
	}

	return c.JSON(200, helper.SuccessDataResponseHelper(msg, CoreToResponse(data)))
}
