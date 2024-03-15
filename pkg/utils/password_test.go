package utils

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHashPassword(t *testing.T) {
	t.Run("Hash password", func(t *testing.T) {
		password := "test*password"
		hashedPassword, err := HashPassword(password)
		assert.NoError(t, err)

		os.Setenv("password", password)
		os.Setenv("hashedPassword", hashedPassword)
	})
}

func TestVerifyPassword(t *testing.T) {
	t.Run("Verify password", func(t *testing.T) {
		password := os.Getenv("password")
		hashedPassword := os.Getenv("hashedPassword")

		err := VerifyPassword(hashedPassword, password)
		assert.NoError(t, err)
		os.Unsetenv("password")
		os.Unsetenv("hashedPassword")
	})
}
