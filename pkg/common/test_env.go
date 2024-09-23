package common

import "os"

func SetUpTestEnv() {
	os.Setenv("GIN_MODE", "debug")
	os.Setenv("TWELVE_FACTOR_MODE", "true")

	os.Setenv("DSN", "host=localhost user=superuser password=youmayneverguess dbname=ginstarter port=5432 sslmode=disable")
	os.Setenv("REDIS_URL", "localhost:6379")

	os.Setenv("JWT_ACCESS_TOKEN_SECRET", "secret**for**access**token")
	os.Setenv("JWT_REFRESH_TOKEN_SECRET", "secret**for**refresh**token")
	os.Setenv("JWT_ACCESS_TOKEN_EXPIRES_IN", "3600s")
	os.Setenv("JWT_REFRESH_TOKEN_EXPIRES_IN", "86400s")
}
