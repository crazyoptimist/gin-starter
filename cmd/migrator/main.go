package main

import (
	"log"

	"gin-starter/internal/config"
	"gin-starter/internal/model"
)

func main() {
	if err := config.LoadConfig(".env"); err != nil {
		log.Fatalln(`Please make sure .env file exists or env variable TWELVE_FACTOR_MODE is set to "true": `, err)
	}

	if err := config.ConnectDB(); err != nil {
		log.Fatalln("Error: database connection failed: ", err)
	}

	if err := config.Config.DB.AutoMigrate(&model.User{}); err != nil {
		log.Fatalln("Database migration failed: ", err)
	}
}
