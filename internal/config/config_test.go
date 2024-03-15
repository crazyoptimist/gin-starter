package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadConfig(t *testing.T) {

	t.Run("Return error when dotenv file is not found AND TWELVE_FACTOR_MODE is set to false", func(t *testing.T) {
		os.Setenv("TWELVE_FACTOR_MODE", "false")

		err := LoadConfig("./notexisting.env")

		assert.NotNil(t, err)
	})

	t.Run("Load config from shell environment when TWELVE_FACTOR_MODE is set to true", func(t *testing.T) {
		os.Setenv("TWELVE_FACTOR_MODE", "true")

		sampleDsn := "host=localhost user=username password=password dbname=gin_starter port=5432 sslmode=disable TimeZone=America/Chicago"

		os.Setenv("DSN", sampleDsn)

		LoadConfig("./notexisting.env")

		assert.Equal(t, sampleDsn, Config.DSN)
	})
}
