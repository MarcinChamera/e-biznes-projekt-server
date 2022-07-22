package database

import (
	"backend/database/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var database *gorm.DB = nil

func Connect() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("ebiznes.db"))
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&models.Product{})
	db.AutoMigrate(&models.Payment{})
	db.AutoMigrate(&models.Order{})
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Category{})
	db.AutoMigrate(&models.Cart{})

	database = db

	return db
}

func GetDatabase() *gorm.DB {
	return database
}
