package util

import (
	"devcode/config"
	"devcode/model"
	"strings"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB(config *config.ServerConfig) *gorm.DB {

	dsnString := []string{
		config.Database.Username, ":", config.Database.Password, "@tcp(", config.Database.Host, ":", config.Database.Port, ")/", config.Database.Name, "?parseTime=true&loc=Asia%2FJakarta&charset=utf8mb4&collation=utf8mb4_unicode_ci",
	}
	dsn := strings.Join(dsnString, "")

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	return db
}

func InitialMigrate(config *config.ServerConfig, db *gorm.DB) {
	if config.Mode == "DEV" {
		db.Migrator().DropTable(&model.Activity{})
		db.Migrator().DropTable(&model.Todo{})
		db.AutoMigrate(&model.Activity{})
		db.AutoMigrate(&model.Todo{})
	} else {
		db.AutoMigrate(&model.Activity{})
		db.AutoMigrate(&model.Todo{})
	}
}
