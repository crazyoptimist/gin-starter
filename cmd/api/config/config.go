package config

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Config appConfig

type appConfig struct {
	DB         *gorm.DB
	ServerPort int    `mapstructure:"server_port"`
	DSN        string `mapstructure:"dsn"`
}

func LoadConfig(configPaths ...string) error {
	v := viper.New()
	v.SetConfigName("api")
	v.SetConfigType("yaml")
	v.AutomaticEnv()

	v.SetDefault("server_port", 8080)
	v.SetDefault("dsn", fmt.Sprintf("%v", v.Get("DSN")))

	for _, path := range configPaths {
		v.AddConfigPath(path)
	}

	if err := v.ReadInConfig(); err != nil {
		return fmt.Errorf("failed to read the configuration file: %s", err)
	}

	return v.Unmarshal(&Config)
}

func ConnectDB() {
	var dbErr error
	Config.DB, dbErr = gorm.Open(postgres.Open(Config.DSN), &gorm.Config{})
	if dbErr != nil {
		log.Fatalln("Database connection failed: ", dbErr)
	}
}
