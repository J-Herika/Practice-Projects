package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type User struct {
	ID           uint   `gorm:"primaryKey;autoIncrement"`
	Username     string `gorm:"size:100;unique"`
	PasswordHash string
}

type UserInfo struct {
	Username string
	Password string
}

var DB *gorm.DB

func main() {
	var err error
	DB, err = gorm.Open(sqlite.Open("Gorm.db"), &gorm.Config{})
	if err != nil {
		log.Fatalln("Couldn't Connect to DataBase: ", err)
	}
	errConn := DB.AutoMigrate(&User{})
	if errConn != nil {
		log.Fatalln("Couldn't Create DataBase schema: ", errConn)
	}

	http.HandleFunc("/register", registerHandler)
	http.HandleFunc("/login", loginHandler)
	errServer := http.ListenAndServe(":3000", nil)
	if errServer != nil {
		fmt.Printf("COULDN'T CONNECT TO SERVER")
	}
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "THIS METHOD ONLY ACCEPTS POST REQUESTS", http.StatusBadRequest)
		return
	}

	var loginData UserInfo
	err := json.NewDecoder(r.Body).Decode(&loginData)
	if err != nil {
		http.Error(w, "COULD NOT DECODE THE LOGIN USER JSON", http.StatusBadRequest)
		return
	}

	var user User
	result := DB.Where("Username = ?", loginData.Username).First(&user)
	if result.Error != nil {
		http.Error(w, "USER NOT FOUND", http.StatusNotFound)
		return
	}
	if !passwordComparer([]byte(user.PasswordHash), loginData.Password) {
		http.Error(w, "WRONG PASSWORD", http.StatusUnauthorized)
		return
	}

	strToken := generateJWT(user.Username)
	if strToken == "" {
		http.Error(w, "COULD NOT FETCH TOKEN", http.StatusInternalServerError)

	}

	if errEncode := json.NewEncoder(w).Encode(map[string]string{"Token": strToken}); errEncode != nil {
		http.Error(w, "COULD NOT FETCH TOKEN", http.StatusInternalServerError)
	}

}

func registerHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only accepts post methods", http.StatusBadRequest)
		return
	}

	var user UserInfo
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "COULD NOT DECODE JSON", http.StatusBadRequest)
		return
	}

	// first we hash the password  ps: never store a password or do anything with it without hashing
	hashedPass := passwordHasher(user.Password)

	// create a new user "Hard coding it for now"
	newUser := User{Username: user.Username, PasswordHash: string(hashedPass)}

	// adding the user to the database here ps: here am using a {&} to change the variable it self and not to create a copy
	if errNewUser := addUser(&newUser); errNewUser != nil {
		http.Error(w, errNewUser.Error(), http.StatusConflict)
		return
	}

	// testing purpose
	w.Header().Set("Content-Type", "application/json")
	errEncode := json.NewEncoder(w).Encode(&newUser)
	if errEncode != nil {
		http.Error(w, "Could NOT SEND USER INFO", http.StatusInternalServerError)
		return
	}
}

// here i am using {*} to access the actual user data instead of creating a copy.
func addUser(user *User) error {
	result := DB.Create(user)
	if result.Error != nil {
		return fmt.Errorf("FAILED TO ADD USER: %w", result.Error)
	}
	return nil
}

func passwordHasher(password string) []byte {

	hashedPass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("Password Hashing failed: ", err)
		return nil
	}
	return hashedPass
}

func passwordComparer(hashed []byte, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(password))
	return err == nil
}

func generateJWT(username string) string {
	claim := jwt.MapClaims{
		"Username": username,
		"exp":      time.Now().Add(time.Hour * 1).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	strToken, err := token.SignedString([]byte("verySecure"))
	if err != nil {
		log.Printf("COULD NOT STRINGIFY TOKEN")
		return ""
	}

	return strToken
}
