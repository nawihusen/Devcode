package main

import (
	"skyshi/config"
	"skyshi/factory"
	"skyshi/migrations"
	"skyshi/utils/database"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	cfg := config.GetConfig()
	db := database.InitDBmySql(cfg)
	migrations.InitMigrate(db)

	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE, echo.PATCH},
	}))

	factory.InitFactory(e, db)
	e.Logger.Fatal(e.Start(":3030"))
}
