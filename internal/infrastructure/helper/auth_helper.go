package helper

import (
	"context"
	"fmt"

	"gin-starter/internal/config"
	"gin-starter/internal/domain/auth"
)

type authHelper struct{}

var _ auth.AuthHelper = (*authHelper)(nil)

func NewAuthHelper() *authHelper {
	return &authHelper{}
}

const BLACKLISTED_TOKEN_PREFIX = "blacklist:"

func (a *authHelper) BlacklistToken(token string) error {
	ctx := context.Background()
	key := fmt.Sprintf("%s:%s", BLACKLISTED_TOKEN_PREFIX, token)
	return config.Global.RedisClient.Set(
		ctx,
		key,
		"true",
		config.Global.JwtRefreshTokenExpiresIn,
	).Err()
}

func (a *authHelper) IsTokenBlacklisted(token string) (bool, error) {
	ctx := context.Background()
	key := fmt.Sprintf("%s:%s", BLACKLISTED_TOKEN_PREFIX, token)
	val, _ := config.Global.RedisClient.Get(ctx, key).Result()

	if val != "" {
		return true, nil
	}

	return false, nil
}
