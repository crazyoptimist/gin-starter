package auth

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateAccessToken(t *testing.T) {
	t.Run("it should generate a valid access token", func(t *testing.T) {
		userId := uint(1)
		accessToken, err := GenerateAccessToken(userId)
		assert.NoError(t, err)

		isValid, _, _, _ := ValidateToken(accessToken)

		if isValid != true {
			t.Errorf("Expected a valid access token, but got an invalid one:\n %v ", accessToken)
		}
	})
}
