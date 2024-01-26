package service_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"gin-starter/internal/repository"
	"gin-starter/internal/service"
)

func TestFindById(t *testing.T) {
	userRepository := repository.NewMockUserRepository()
	userService := service.NewUserService(userRepository)

	t.Run("it should return error for not existing id", func(t *testing.T) {
		_, err := userService.FindById(100)
		assert.NotNil(t, err)
	})
}
