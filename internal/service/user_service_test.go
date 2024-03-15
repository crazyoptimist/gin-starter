package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindById(t *testing.T) {
	userRepository := NewUserRepositoryStub()
	userService := NewUserService(userRepository)

	t.Run("it should return error for not existing id", func(t *testing.T) {
		_, err := userService.FindById(100)
		assert.NotNil(t, err)
	})
}
