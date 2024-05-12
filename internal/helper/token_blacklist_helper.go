package helper

import (
	"context"
	"fmt"

	"gin-starter/internal/config"
)

const BLACKLISTED_TOKEN_PREFIX = "blacklist:"

func BlacklistToken(token string) error {
	ctx := context.Background()
	key := fmt.Sprintf("%s:%s", BLACKLISTED_TOKEN_PREFIX, token)
	return config.Config.RedisClient.Set(
		ctx,
		key,
		"true",
		config.Config.JwtRefreshTokenExpiresIn,
	).Err()
}

func IsTokenBlacklisted(token string) (bool, error) {
	ctx := context.Background()
	key := fmt.Sprintf("%s:%s", BLACKLISTED_TOKEN_PREFIX, token)
	val, _ := config.Config.RedisClient.Get(ctx, key).Result()

	if val != "" {
		return true, nil
	}

	return false, nil
}
