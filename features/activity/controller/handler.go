package controller

import (
	"fmt"
	"skyshi/features/activity"
	"skyshi/utils/helper"
	"strconv"

	"github.com/labstack/echo/v4"
)

type Controller struct {
	control activity.ServiceInterface
}

func New(e *echo.Echo, data activity.ServiceInterface) {
	handler := &Controller{
		control: data,
	}
	e.GET("/activity-groups", handler.Activities)
	e.GET("/activity-groups/:id", handler.Activity)
	e.POST("/activity-groups", handler.Create)
	e.DELETE("/activity-groups/:id", handler.Delete)
	e.PATCH("/activity-groups/:id", handler.Update)

}

func (client *Controller) Activities(c echo.Context) error {
	data, msg, err := client.control.GetActivities()
	if err != nil {
		return c.JSON(400, helper.FailedResponseHelper(msg))
	}

	return c.JSON(200, helper.SuccessDataResponseHelper(msg, CoreToResponseList(data)))
}

func (client *Controller) Activity(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(400, helper.FailedResponseHelper("Parameter must be number"))
	}

	data, msg, er := client.control.GetActivity(uint(id))
	if er != nil {
		return c.JSON(400, helper.FailedResponseHelper(msg))
	} else if msg == "Not Found" {
		notfound := fmt.Sprintf("Activity with ID %d Not Found", id)
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
		return c.JSON(400, helper.BadRequest())
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
		notfound := fmt.Sprintf("Activity with ID %d Not Found", id)
		return c.JSON(404, helper.NotFoundHelper(notfound))
	}

	return c.JSON(200, helper.SuccessResponseHelper(msg))
}

func (client *Controller) Update(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(400, helper.FailedResponseHelper("Parameter must be number"))
	}

	var req Request
	er := c.Bind(&req)
	if er != nil {
		return c.JSON(400, helper.FailedResponseHelper("Failed Bind"))
	} else if req.Title == "" && req.Email == "" {
		return c.JSON(400, helper.BadRequest())
	}

	data, msg, er := client.control.Update(uint(id), req.ToCore())
	if er != nil {
		return c.JSON(400, helper.FailedResponseHelper(msg))
	} else if msg == "Not Found" {
		notfound := fmt.Sprintf("Activity with ID %d Not Found", id)
		return c.JSON(404, helper.NotFoundHelper(notfound))
	}

	return c.JSON(200, helper.SuccessDataResponseHelper(msg, CoreToResponse(data)))
}
