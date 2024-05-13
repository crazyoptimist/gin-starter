package config

import (
	"context"
	"os"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Config appConfig

type appConfig struct {
	DB          *gorm.DB
	RedisClient *redis.Client

	ServerPort               int           `mapstructure:"SERVER_PORT"`
	DSN                      string        `mapstructure:"DSN"`
	RedisUrl                 string        `mapstructure:"REDIS_URL"`
	JwtAccessTokenSecret     string        `mapstructure:"JWT_ACCESS_TOKEN_SECRET"`
	JwtRefreshTokenSecret    string        `mapstructure:"JWT_REFRESH_TOKEN_SECRET"`
	JwtAccessTokenExpiresIn  time.Duration `mapstructure:"JWT_ACCESS_TOKEN_EXPIRES_IN"`
	JwtRefreshTokenExpiresIn time.Duration `mapstructure:"JWT_REFRESH_TOKEN_EXPIRES_IN"`
}

func LoadConfig(cfgFile string) error {
	v := viper.New()

	v.SetDefault("SERVER_PORT", 8080)

	// Enable twelve-factor methodology
	if os.Getenv("TWELVE_FACTOR_MODE") == "true" {
		v.AutomaticEnv()
		v.BindEnv("DSN")
		v.BindEnv("REDIS_URL")
		v.BindEnv("JWT_ACCESS_TOKEN_SECRET")
		v.BindEnv("JWT_REFRESH_TOKEN_SECRET")
		v.BindEnv("JWT_ACCESS_TOKEN_EXPIRES_IN")
		v.BindEnv("JWT_REFRESH_TOKEN_EXPIRES_IN")

		return v.Unmarshal(&Config)
	}

	v.SetConfigType("env")
	v.SetConfigFile(cfgFile)

	if err := v.ReadInConfig(); err != nil {
		return err
	}

	return v.Unmarshal(&Config)
}

func ConnectDB() error {
	db, err := gorm.Open(postgres.Open(Config.DSN), &gorm.Config{})
	if err != nil {
		return err
	}

	Config.DB = db
	return nil
}

func ConnectRedis() error {
	rdb := redis.NewClient(&redis.Options{Addr: Config.RedisUrl})

	ctx := context.Background()
	err := rdb.Set(ctx, "ping", "pong", time.Second).Err()
	if err != nil {
		return err
	}

	Config.RedisClient = rdb
	return nil
}
