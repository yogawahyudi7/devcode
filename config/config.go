package config

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/gommon/log"
)

type ServerConfig struct {
	Mode     string
	Database struct {
		Name     string
		Host     string
		Port     string
		Username string
		Password string
	}
}

func Get() *ServerConfig {

	err := godotenv.Load()
	if err != nil {
		log.Info("Error loading .env file")
	}

	var defaultConfig ServerConfig

	defaultConfig.Mode = os.Getenv("MODE")
	defaultConfig.Database.Host = os.Getenv("MYSQL_HOST")
	defaultConfig.Database.Port = os.Getenv("MYSQL_PORT")
	defaultConfig.Database.Username = os.Getenv("MYSQL_USER")
	defaultConfig.Database.Password = os.Getenv("MYSQL_PASSWORD")
	defaultConfig.Database.Name = os.Getenv("MYSQL_DBNAME")

	return &defaultConfig
}
