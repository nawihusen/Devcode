package factory

import (
	activityController "skyshi/features/activity/controller"
	activityData "skyshi/features/activity/data"
	activitySetvice "skyshi/features/activity/service"

	todoController "skyshi/features/todo/controller"
	todoData "skyshi/features/todo/data"
	todoService "skyshi/features/todo/service"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitFactory(e *echo.Echo, db *gorm.DB) {
	activityDataFactory := activityData.New(db)
	activityUsecaseFactory := activitySetvice.New(activityDataFactory)
	activityController.New(e, activityUsecaseFactory)

	todoDataFactory := todoData.New(db)
	todoUsecaseFactory := todoService.New(todoDataFactory)
	todoController.New(e, todoUsecaseFactory)
}
