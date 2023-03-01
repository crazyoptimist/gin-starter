package main

import (
	"log"

	"gin-starter/internal/core/config"
	"gin-starter/internal/user"
)

func main() {
	if err := config.LoadConfig(".env"); err != nil {
		log.Println("Warning: dotenv file is missing, please make sure you have configured environment variables properly", err)
	}

	if err := config.ConnectDB(); err != nil {
		log.Fatalln("Error: database connection failed: ", err)
	}

	if err := config.Config.DB.AutoMigrate(&user.User{}); err != nil {
		log.Fatalln("Database migration failed: ", err)
	}
}
