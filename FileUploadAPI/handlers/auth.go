package handlers

import (
	"FileUploadAPI/database"
	"FileUploadAPI/utils"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

func UploadFileHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Wrong Method send. accept POST Method only", http.StatusBadRequest)
		return
	}

	// tells go to parse the data("Break it up into different parts like file,metadata etc.") then tells it to accept only file sizes up to 10mb.
	errParse := r.ParseMultipartForm(10 << 20)
	if errParse != nil {
		http.Error(w, "Could not Parse Data.", http.StatusBadRequest)
		return
	}

	// extracts the uploaded file from the data. also extracts the metadata("handler").
	file, handler, errFile := r.FormFile("file")
	if errFile != nil {
		http.Error(w, "Could not Get File. ", http.StatusBadRequest)
		return
	}

	// Create the file in the "upload" dir with the name of the file.                ps: if the file already exists i does not create it again
	dst, errCreate := os.Create("./uploads/" + handler.Filename)
	if errCreate != nil {
		http.Error(w, "Error Saving file", http.StatusInternalServerError)
		return
	}

	_, errCopy := io.Copy(dst, file)
	if errCopy != nil {
		http.Error(w, "Error Copying file content", http.StatusInternalServerError)
		return
	}

	fmt.Printf("Uploaded successfully")
	defer file.Close()
}

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

	errEncode := json.NewEncoder(w).Encode(map[string]string{"Token": strToken})
	if errEncode != nil {
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
