package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadConfig(t *testing.T) {
	t.Run("Load config from shell environment when TWELVE_FACTOR_MODE is set to true", func(t *testing.T) {
		os.Setenv("TWELVE_FACTOR_MODE", "true")

		sampleDSN := "host=localhost user=username password=password dbname=gin_starter port=5432 sslmode=disable TimeZone=America/Chicago"

		os.Setenv("DSN", sampleDSN)

		if err := LoadConfig(""); err != nil {
			t.Errorf("Loading config failed: %v", err)
		}

		assert.Equal(t, sampleDSN, Config.DSN)
	})
}
