package helper

import (
	"errors"
	"fmt"

	badger "github.com/dgraph-io/badger/v4"

	"gin-starter/internal/config"
)

const BLACKLISTED_TOKEN_PREFIX = "bl"

func BlacklistToken(token string) error {
	key := fmt.Sprintf("%s:%s", BLACKLISTED_TOKEN_PREFIX, token)
	return SetCacheValue(key, "true", config.Config.JwtRefreshTokenExpiresIn)
}

func IsTokenBlacklisted(token string) (bool, error) {
	key := fmt.Sprintf("%s:%s", BLACKLISTED_TOKEN_PREFIX, token)
	val, err := GetCacheValue(key)
	if errors.Is(err, badger.ErrKeyNotFound) {
		return false, nil
	}

	if val != "" {
		return true, nil
	}

	return false, nil
}
