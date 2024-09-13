package initialisers

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// var DB *gorm.DB

// UTIL --> FUNCTIONS
func ConnectToDB() *gorm.DB {
	fmt.Println("Connecting to database...")

	DB, err := gorm.Open(sqlite.Open("books.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to Database: " + err.Error())
	}
	fmt.Println("Connected to database")

	return DB
}
