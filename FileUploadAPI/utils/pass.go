package utils

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

func PasswordHasher(password string) []byte {

	hashedPass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Print("Password Hashing failed: ", err)
		return nil
	}
	return hashedPass
}

func PasswordComparer(hashed []byte, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(password))
	return err != nil
}
