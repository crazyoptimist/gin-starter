package main

import (
	"fmt"
	"gin-starter/internal/core/config"
	"gin-starter/internal/user"
)

func main() {
	if err := config.LoadConfig(".env"); err != nil {
		panic(fmt.Errorf("invalid application configuration: %s", err))
	}
	config.ConnectDB()
	config.Config.DB.AutoMigrate(&user.User{})
}
