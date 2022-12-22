package config

import (
	"os"
	"sync"
)

type AppConfig struct {
	MYSQL_HOST     string
	MYSQL_PORT     string
	MYSQL_USER     string
	MYSQL_PASSWORD string
	MYSQL_DBNAME   string
}

var lock = &sync.Mutex{}
var config *AppConfig

func GetConfig() *AppConfig {

	lock.Lock()
	defer lock.Unlock()

	if config == nil {
		config = initConfig()
	}
	return config

}

func initConfig() *AppConfig {

	var defaultConfig AppConfig
	defaultConfig.MYSQL_HOST = os.Getenv("MYSQL_HOST")
	defaultConfig.MYSQL_PORT = "3306"
	defaultConfig.MYSQL_USER = os.Getenv("MYSQL_USER")
	defaultConfig.MYSQL_PASSWORD = os.Getenv("MYSQL_PASSWORD")
	defaultConfig.MYSQL_DBNAME = os.Getenv("MYSQL_DBNAME")
	return &defaultConfig
}
