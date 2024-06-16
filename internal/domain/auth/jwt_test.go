package auth

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"

	"gin-starter/internal/config"
)

func TestGenerateAccessToken(t *testing.T) {
	workdir, _ := os.Getwd()
	envpath := "../../../.env"
	_ = config.LoadConfig(filepath.Join(workdir, envpath))

	t.Run("Generate a valid access token", func(t *testing.T) {

		userId := 1
		accessToken, err := GenerateAccessToken(userId)
		assert.NoError(t, err)

		isValid, _, _, _ := ValidateToken(accessToken)

		if isValid != true {
			t.Errorf("Expected a valid access token, but got an invalid one:\n %v ", accessToken)
		}
	})
}
