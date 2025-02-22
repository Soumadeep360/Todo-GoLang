package database

import (
	"fmt"
	"log"
	"todo-app/config"
	"todo-app/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Could not load config", err)
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database", err)
	}

	db.AutoMigrate(&models.Todo{}) // Migrate the Todo model

	DB = db
}
