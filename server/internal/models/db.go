package models

import (
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB
var err error

func NewMySQLClient() {
	db, err = gorm.Open(mysql.Open(os.Getenv("DB_CONNECTION_STRING")), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	if err = db.AutoMigrate(&Task{}); err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&Task{}, &User{})
}
