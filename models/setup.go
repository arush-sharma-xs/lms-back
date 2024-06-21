package models

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {

	db, err := gorm.Open(sqlite.Open("libraryMS.db"), &gorm.Config{})

	if err != nil {
		panic("Failed to Connected")
	}

	db.AutoMigrate(&Library{}, &Users{}, &BookInventory{}, &RequestEvents{}, &IssueRegistery{})

	DB = db
}
