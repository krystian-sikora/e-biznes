package config

import (
	"golang/src/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var GormDB *gorm.DB

func ConnectDatabase() {
	db, err := gorm.Open(sqlite.Open("/app/data/database.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	err = db.AutoMigrate(&model.Product{})
	if err != nil {
		panic("failed to migrate database")
	}
	err = db.AutoMigrate(&model.Category{})
	if err != nil {
		panic("failed to migrate database")
	}
	err = db.AutoMigrate(&model.Basket{})
	if err != nil {
		panic("failed to migrate database")
	}
	err = db.AutoMigrate(&model.Payment{})
	if err != nil {
		panic("failed to migrate database")
	}

	GormDB = db
}
