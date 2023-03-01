package config

import (
	"fmt"
	"time"

	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Config appConfig

type appConfig struct {
	DB *gorm.DB

	ServerPort int    `mapstructure:"SERVER_PORT"`
	DSN        string `mapstructure:"DSN"`

	JwtAccessTokenSecret     string        `mapstructure:"JWT_ACCESS_TOKEN_SECRET"`
	JwtRefreshTokenSecret    string        `mapstructure:"JWT_REFRESH_TOKEN_SECRET"`
	JwtAccessTokenExpiresIn  time.Duration `mapstructure:"JWT_ACCESS_TOKEN__EXPIRES_IN"`
	JwtRefreshTokenExpiresIn time.Duration `mapstructure:"JWT_REFRESH_TOKEN_EXPIRES_IN"`
}

func LoadConfig(configFile string) error {
	v := viper.New()
	v.SetConfigFile(configFile)
	v.AutomaticEnv()

	v.SetDefault("SERVER_PORT", 8080)

	if err := v.ReadInConfig(); err != nil {
		return fmt.Errorf("failed to read the configuration file: %s", err)
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
