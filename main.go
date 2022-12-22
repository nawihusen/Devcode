package main

import (
	"fmt"
	"skyshi/config"
	"skyshi/factory"
	"skyshi/migrations"
	"skyshi/utils/database"

	"github.com/labstack/echo/v4"
)

func main() {
	cfg := config.GetConfig()
	db := database.InitDBmySql(cfg)
	migrations.InitMigrate(db)

	e := echo.New()

	factory.InitFactory(e, db)
	fmt.Println("test")
	e.Logger.Fatal(e.Start(":3030"))
}
