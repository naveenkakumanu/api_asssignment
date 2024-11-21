package config

import (
	"api_assignment/models"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite" // SQLite
)

var DB *gorm.DB

func Connect() {
	var err error
	DB, err = gorm.Open("sqlite3", "./gorm.db")
	if err != nil {
		log.Fatal("failed to connect to database")
	}
	DB.AutoMigrate(&models.User{})
}
