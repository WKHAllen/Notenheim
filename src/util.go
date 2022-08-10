package src

import (
	"math/rand"
	"time"

	"golang.org/x/crypto/bcrypt"
)

const cost = 12
const b64chars = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789-_"

func init() {
    rand.Seed(time.Now().UnixNano())
}

// Base64ID generates a new random ID in base 64
func Base64ID(length int) string {
	id := make([]byte, length)
	for i := range id {
		id[i] = b64chars[rand.Intn(len(b64chars))]
	}
	return string(id)
}

// GetTime returns the current time in 
func GetTime() int64 {
	return time.Now().Unix()
}

// HashPassword generates a password hashed using bcrypt
func HashPassword(password string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	return string(hashed), err
}

// ComparePassword compares a hashed password to a given password
func ComparePassword(password string, hashedPassword string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)) == nil
}
