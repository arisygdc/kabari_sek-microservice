package user

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

const layout string = "2006-06-02"

// Check hash using bcrypt
func CheckHashPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

// Generate bcrypt hash password, return hashed password
func HashPassword(password string) (string, error) {
	hashValue, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashValue), err
}

// timeToParse format date "2006-Jan-02"
func ParseTime(timeToParse string) (time.Time, error) {
	return time.Parse(layout, timeToParse)
}
