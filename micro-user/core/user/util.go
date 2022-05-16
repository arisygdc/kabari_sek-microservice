package user

import "golang.org/x/crypto/bcrypt"

// Check hash using bcrypt
func CheckHashPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

// Generate bcrypt hash password, return hashed password
func HashPassword(password string) (string, error) {
	hashValue, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashValue), err
}
