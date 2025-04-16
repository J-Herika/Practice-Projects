package main

import (
	"FileUploadAPI/database"
	"FileUploadAPI/handlers"
	"fmt"
	"net/http"

	"gorm.io/gorm"
)

var DB *gorm.DB

func main() {
	database.Init()

	http.HandleFunc("/register", handlers.RegisterHandler)
	http.HandleFunc("/login", handlers.LoginHandler)
	http.HandleFunc("/upload", handlers.UploadFileHandler)
	errServer := http.ListenAndServe(":3000", nil)
	if errServer != nil {
		fmt.Printf("COULDN'T CONNECT TO SERVER")
	}
}
