package helper

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"gin-starter/internal/config"
)

func TestGenerateAccessToken(t *testing.T) {
	currentDirPath, _ := os.Getwd()
	relativeEnvPath := "../../.env.test"

	if err := config.LoadConfig(currentDirPath + "/" + relativeEnvPath); err != nil {
		panic("Failed to load test env file")
	}

	t.Run("Generate a valid access token", func(t *testing.T) {

		userId := uint(1)
		accessToken, err := GenerateAccessToken(userId)
		assert.NoError(t, err)

		isValid, _, _, _ := ValidateToken(accessToken)

		if isValid != true {
			t.Errorf("Expected a valid access token, but got an invalid one:\n %v ", accessToken)
		}
	})
}
