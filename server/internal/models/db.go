package models

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB
var err error

func NewMySQLClient() {
	conn := "todo:Password123!@(localhost:3306)/todoapp?charset=utf8&parseTime=True&loc=Local"

	db, err = gorm.Open(mysql.Open(conn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	if err = db.AutoMigrate(&Task{}); err != nil {
		log.Fatal(err)
	}
}
