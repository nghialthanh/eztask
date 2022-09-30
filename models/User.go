package models

import (
	"fmt"
	"log"
	"time"

	"golang.org/x/crypto/bcrypt"
)

// User is the model that governs all notes objects retrived or inserted into the DB
type User struct {
	ID            string    `json:"id" gorm:"primaryKey"`
	First_name    string    `json:"first_name" gorm:"size:255;not null"`
	Last_name     string    `json:"last_name" gorm:"size:255;not null"`
	Password      string    `json:"Password" gorm:"size:30;not null"`
	Email         string    `json:"email" gorm:"size:255;not null;unique"`
	Phone         string    `json:"phone" gorm:"size:20;not null;unique"`
	Token         string    `json:"token"`
	User_type     string    `json:"user_type" gorm:"size:30;not null"`
	Refresh_token string    `json:"refresh_token"`
	Created_at    time.Time `json:"created_at"`
	Updated_at    time.Time `json:"updated_at"`
	User_id       string    `json:"user_id"`
	// Project       []RoleProject      `json:"project"`
}

// HashPassword is used to encrypt the password before it is stored in the DB
func HashPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		log.Panic(err)
	}

	return string(bytes)
}

// VerifyPassword checks the input password while verifying it with the passward in the DB.
func VerifyPassword(userPassword string, providedPassword string) (bool, string) {
	err := bcrypt.CompareHashAndPassword([]byte(providedPassword), []byte(userPassword))
	check := true
	msg := ""

	if err != nil {
		msg = fmt.Sprintf("login or passowrd is incorrect")
		check = false
	}

	return check, msg
}
