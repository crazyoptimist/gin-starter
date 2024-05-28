package config

import (
	"context"
	"errors"
	"os"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Global AppConfig

type AppConfig struct {
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

func (c *AppConfig) Validate() error {
	if c.DSN == "" ||
		c.RedisUrl == "" ||
		len(c.JwtAccessTokenSecret) < 16 ||
		len(c.JwtRefreshTokenSecret) < 16 ||
		c.JwtAccessTokenExpiresIn == 0 ||
		c.JwtRefreshTokenExpiresIn == 0 {
		return errors.New("Environment variables validation failed.")
	}
	return nil
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

		if err := v.Unmarshal(&Global); err != nil {
			return err
		}

		return Global.Validate()
	}

	v.SetConfigType("env")
	v.SetConfigFile(cfgFile)

	if err := v.ReadInConfig(); err != nil {
		return err
	}

	if err := v.Unmarshal(&Global); err != nil {
		return err
	}

	return Global.Validate()
}

func ConnectDB() error {
	db, err := gorm.Open(postgres.Open(Global.DSN), &gorm.Config{})
	if err != nil {
		return err
	}

	Global.DB = db
	return nil
}

func ConnectRedis() error {
	rdb := redis.NewClient(&redis.Options{Addr: Global.RedisUrl})

	ctx := context.Background()
	err := rdb.Set(ctx, "ping", "pong", time.Second).Err()
	if err != nil {
		return err
	}

	Global.RedisClient = rdb
	return nil
}
