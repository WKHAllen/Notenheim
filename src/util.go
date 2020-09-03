package src

import (
	"golang.org/x/crypto/bcrypt"
	"math/rand"
	"time"
)

const cost = 12
const b64chars = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789-_"

func init() {
    rand.Seed(time.Now().UnixNano())
}

// Base64ID generates a new random ID in base 64
func Base64ID(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = b64chars[rand.Intn(len(b64chars))]
	}
	return string(b)
}

// GetTime returns the current time in 
func GetTime() int64 {
	return time.Now().Unix()
}

// HashPassword generates a password hashed using bcrypt
func HashPassword(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), cost)
}

// ComparePassword compares a hashed password to a given password
func ComparePassword(password string, hashedPassword []byte) bool {
	return bcrypt.CompareHashAndPassword(hashedPassword, []byte(password)) == nil
}
