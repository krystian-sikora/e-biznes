package config

import (
	"fmt"
	"golang/src/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var GormDB *gorm.DB

const ConnectionFailure = "failed to connect database"

func ConnectDatabase() {
	fmt.Println("connecting to test db")
	db, err := gorm.Open(sqlite.Open("/app/data/test.db"), &gorm.Config{})
	if err != nil {
		panic(ConnectionFailure)
	}
	err = db.AutoMigrate(&model.Product{})
	if err != nil {
		panic(ConnectionFailure)
	}
	err = db.AutoMigrate(&model.Category{})
	if err != nil {
		panic(ConnectionFailure)
	}
	err = db.AutoMigrate(&model.Basket{})
	if err != nil {
		panic(ConnectionFailure)
	}
	err = db.AutoMigrate(&model.Payment{})
	if err != nil {
		panic(ConnectionFailure)
	}

	GormDB = db
}
