package user

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindById(t *testing.T) {
	userRepository := NewUserRepositoryStub()
	userService := NewUserService(userRepository)

	t.Run("Return error for not existing user ID", func(t *testing.T) {
		_, err := userService.FindById(100)
		assert.NotNil(t, err)
	})
}
