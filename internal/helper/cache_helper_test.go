package helper

import (
	"testing"
	"time"

	"github.com/dgraph-io/badger/v4"
	"github.com/stretchr/testify/assert"

	"gin-starter/internal/config"
)

func TestSetCacheValue(t *testing.T) {
	if err := config.ConnectCacheDB(); err != nil {
		t.Error("Cache DB connection failed: ", err)
	}

	t.Run("Sets a value in the cache store", func(t *testing.T) {
		err := SetCacheValue("ping", "pong", 10*time.Hour)
		if err != nil {
			t.Error(err)
		}
	})
}

func TestGetCacheValue(t *testing.T) {
	defer config.Config.CacheClient.Close()

	t.Run("Gets a value from the cache store", func(t *testing.T) {
		val, err := GetCacheValue("ping")
		if err != nil {
			t.Error(err)
		}

		assert.Equal(t, "pong", val)
	})

	t.Run("Can't get an expired value from the cache store", func(t *testing.T) {
		_ = SetCacheValue("short-lived", "gone", time.Millisecond)

		time.Sleep(time.Millisecond)

		val, err := GetCacheValue("short-lived")

		assert.Equal(t, "", val)
		assert.ErrorIs(t, err, badger.ErrKeyNotFound)
	})
}
