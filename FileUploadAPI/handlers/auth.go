package handlers

import (
	"FileUploadAPI/database"
	"FileUploadAPI/utils"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

func DeleteFileHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Error Accept only delete Method", http.StatusBadRequest)
		return
	}

	_, err := ExtractUserIDFromRequest(r)
	if err != nil {
		// return 401 unauthorized
		http.Error(w, "Error unauthorized", http.StatusUnauthorized)
		return
	}

	// get imgName param value from url
	fileName := r.URL.Query().Get("fileName")
	if fileName == "" {
		http.Error(w, "Error Empty value", http.StatusBadRequest)
		return
	}

	filePath := "./uploads/" + fileName
	err := os.Remove(filePath)

	if err != nil {
		http.Error(w, "Error Could not remove file. Make sure its the right file name OR it's stored.", http.StatusNotFound)
		return
	}

	log.Print("successes!")
}

func DownloadFileHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Error Accept Get Methods only.", http.StatusBadRequest)
		return
	}

	_, err := ExtractUserIDFromRequest(r)
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	// get the name of the image from the url we basically cut the string and get the name only.
	fileName := r.URL.Path[len("/download/"):]
	if fileName == "" {
		http.Error(w, "File is missing.", http.StatusBadRequest)
		return
	}

	filePath := "./uploads/" + fileName
	_, errFile := os.Stat(filePath)
	if errFile != nil {
		http.Error(w, "File Not found.", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Disposition", "attachment; filename="+fileName)
	http.ServeFile(w, r, filePath)

}

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

func ExtractUserIDFromRequest(r *http.Request) (uint, error) {
	// Step 1: Get token string from header
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return 0, errors.New("missing auth header")
	}

	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		return 0, errors.New("invalid auth format")
	}

	tokenString := parts[1]

	// Step 2: Parse the token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method")
		}
		return []byte("your_secret_key"), nil
	})

	if err != nil || !token.Valid {
		return 0, errors.New("invalid token")
	}

	// Step 3: Extract user_id from claims
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return 0, errors.New("could not parse claims")
	}

	userIDFloat, ok := claims["user_id"].(float64) // JWT stores numbers as float64
	if !ok {
		return 0, errors.New("user_id not found")
	}

	return uint(userIDFloat), nil
}
