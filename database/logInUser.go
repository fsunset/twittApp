package database

import (
	"log"

	"github.com/fsunset/twittApp/models"
	"golang.org/x/crypto/bcrypt"
)

// LoginUser logs user in using email/password
func LoginUser(email, password string) (models.User, bool) {

	// Checks if user exists
	usr, usrExists, _ := CheckExistentUser(email)
	if !usrExists {
		return usr, false
	}

	// Compare request-password against DB-password
	requestPassword := []byte(password)
	DBPassword := []byte(usr.Password)

	err := bcrypt.CompareHashAndPassword(DBPassword, requestPassword)
	if err != nil {
		log.Fatal("Invalid Credentials")
		return usr, false
	}

	return usr, true
}
