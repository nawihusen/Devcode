package database

import (
	"fmt"
	"log"
	"skyshi/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDBmySql(config *config.AppConfig) *gorm.DB {

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		config.MYSQL_USER,
		config.MYSQL_PASSWORD,
		config.MYSQL_HOST,
		config.MYSQL_PORT,
		config.MYSQL_DBNAME,
	)
	db, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
	}
	return db

}
