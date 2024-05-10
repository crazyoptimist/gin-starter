package utils

import (
	"golang.org/x/crypto/bcrypt"
)

const HASH_COST_FACTOR = 12

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), HASH_COST_FACTOR)
	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}

// VerifyPassword compares a bcrypt hashed password with its possible plaintext equivalent.
// Returns nil on success, or an error on failure.
func VerifyPassword(hashedPassword, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return err
	}

	return nil
}
