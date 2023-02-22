package utils

import (
	"os"
	"testing"
)

func TestHashPassword(t *testing.T) {
	password := "test*password"
	hashedPassword, err := HashPassword(password)
	if err != nil {
		t.Errorf("Expected no error but got: %v", err)
	}

	os.Setenv("password", password)
	os.Setenv("hashedPassword", hashedPassword)
}

func TestVerifyPassword(t *testing.T) {
	password := os.Getenv("password")
	hashedPassword := os.Getenv("hashedPassword")

	err := VerifyPassword(hashedPassword, password)
	if err != nil {
		t.Errorf("Expected no error but got: %v", err)
	}
}
