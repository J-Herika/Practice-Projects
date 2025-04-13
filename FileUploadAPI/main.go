package main

import (
	"fmt"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type User struct {
	ID           uint   `gorm:"primaryKey"`
	Username     string `gorm:"size:100"`
	PasswordHash string
}

func main() {

	db, err := gorm.Open(sqlite.Open("Gorm.db"), &gorm.Config{})
	if err != nil {
		log.Fatalln("Couldn't Connect to DataBase: ", err)
		return
	}
	errConn := db.AutoMigrate(&User{})
	if errConn != nil {
		log.Fatalln("Couldn't Create DataBase: ", errConn)

	}

	// user1 := User{ID: 1, Username: "James Watson", PasswordHash: "fd023ued44sf444ew4r4w328422"}
	//user2 := User{ID: 2, Username: "Astra Eikon", PasswordHash: "fd023ued44sf444ew4r4w328422"}
	// user3 := User{ID: 3, Username: "Yami Sukarno", PasswordHash: "fd023ued44sf444ew4r4w328422"}

	// db.Create(&user1)
	// db.Create(&user2)
	// db.Create(&user3)

	// Fetching the user again to test

	var fetchedUsers []User
	result := db.Where("ID <= ?", 3).Find(&fetchedUsers)

	if result.Error != nil {
		fmt.Println("DB Error:", result.Error)
	} else {
		fmt.Printf("Fetched Users: %+v\n", fetchedUsers)
	}
}
