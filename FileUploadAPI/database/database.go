package database

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() {

	var err error
	DB, err = gorm.Open(sqlite.Open("Gorm.db"), &gorm.Config{})
	if err != nil {
		log.Fatalln("Couldn't Connect to DataBase: ", err)
	}
	errConn := DB.AutoMigrate(&User{})
	if errConn != nil {
		log.Fatalln("Couldn't Create DataBase schema: ", errConn)
	}

}
