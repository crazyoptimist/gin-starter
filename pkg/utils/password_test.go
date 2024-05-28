package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHashPassword(t *testing.T) {
	t.Run("Hash password", func(t *testing.T) {
		password := "test*password"
		_, err := HashPassword(password)

		assert.NoError(t, err)
	})
}

func TestVerifyPassword(t *testing.T) {
	t.Run("Verify password", func(t *testing.T) {
		password := "test*password"
		hashedPassword, _ := HashPassword(password)

		err := VerifyPassword(hashedPassword, password)

		assert.NoError(t, err)
	})
}
