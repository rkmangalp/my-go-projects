package config

import (
	"log"

	"github.com/jinzhu/gorm"

	_ "github.com/go-sql-driver/mysql"
)

var (
	db *gorm.DB
)

// Connect initializes the database connection
func Connect() {
	d, err := gorm.Open("mysql", "root:new_password@tcp(127.0.0.1:3306)/bookstore?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Fatalf("could not connect to database %v", err)
	}
	db = d
}

// GetDB returns the database connection
func GetDB() *gorm.DB {
	return db
}
