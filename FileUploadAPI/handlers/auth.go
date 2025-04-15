package handlers

import (
	"FileUploadAPI/database"
	"FileUploadAPI/utils"
	"encoding/json"
	"fmt"
	"net/http"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "THIS METHOD ONLY ACCEPTS POST REQUESTS", http.StatusBadRequest)
		return
	}

	var loginData database.UserInfo
	err := json.NewDecoder(r.Body).Decode(&loginData)
	if err != nil {
		http.Error(w, "COULD NOT DECODE THE LOGIN USER JSON", http.StatusBadRequest)
		return
	}

	var user database.User
	result := database.DB.Where("Username = ?", loginData.Username).First(&user)
	if result.Error != nil {
		http.Error(w, "USER NOT FOUND", http.StatusNotFound)
		return
	}
	if !utils.PasswordComparer([]byte(user.PasswordHash), loginData.Password) {
		http.Error(w, "WRONG PASSWORD", http.StatusUnauthorized)
		return
	}

	strToken := utils.GenerateJWT(user.Username)
	if strToken == "" {
		http.Error(w, "COULD NOT FETCH TOKEN", http.StatusInternalServerError)

	}

	if errEncode := json.NewEncoder(w).Encode(map[string]string{"Token": strToken}); errEncode != nil {
		http.Error(w, "COULD NOT FETCH TOKEN", http.StatusInternalServerError)
	}

}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only accepts post methods", http.StatusBadRequest)
		return
	}

	var user database.UserInfo
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "COULD NOT DECODE JSON", http.StatusBadRequest)
		return
	}

	// first we hash the password  ps: never store a password or do anything with it without hashing
	hashedPass := utils.PasswordHasher(user.Password)

	// create a new user "Hard coding it for now"
	newUser := database.User{Username: user.Username, PasswordHash: string(hashedPass)}

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
func addUser(user *database.User) error {
	result := database.DB.Create(user)
	if result.Error != nil {
		return fmt.Errorf("FAILED TO ADD USER: %w", result.Error)
	}
	return nil
}
