package model

import (
	"log"

	"github.com/jonwzh/blog-sys/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

// SetDB func
func SetDB(database *gorm.DB) {
	db = database
}

// ConnectToDB func
func ConnectToDB() *gorm.DB {
	connectingStr := config.GetMysqlConnectingString()
	log.Println("Connecting to db...")
	db, err := gorm.Open("mysql", connectingStr)
	if err != nil {
		panic("Failed to connect database, error: " + err.Error())
	}
	db.SingularTable(true)
	return db
}