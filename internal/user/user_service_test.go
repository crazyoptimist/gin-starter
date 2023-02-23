package user

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUserService(t *testing.T) {
	repository := newMockUserRepository()
	s := NewUserService(repository)
	assert.Equal(t, repository, s.UserRepository)
}

func TestFindById(t *testing.T) {
	userService := NewUserService(newMockUserRepository())

	t.Run("it should return error for not existing id", func(t *testing.T) {
		_, err := userService.FindById(100)
		assert.NotNil(t, err)
	})
}
