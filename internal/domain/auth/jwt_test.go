package auth

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"gin-starter/internal/config"
	"gin-starter/pkg/common"
)

func TestGenerateAccessToken(t *testing.T) {
	common.SetUpTestEnv()
	_ = config.LoadConfig("")

	t.Run("Generate a valid access token", func(t *testing.T) {

		userId := 1
		accessToken, err := GenerateJwtToken(AccessTokenKeyId, userId)
		assert.NoError(t, err)

		isValid, _, _, _ := ValidateJwtToken(accessToken)

		if isValid != true {
			t.Errorf(
				"Expected a valid access token, but got an invalid one:\n %v ",
				accessToken,
			)
		}
	})
}
