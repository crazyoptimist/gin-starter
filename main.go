package main

import (
	"fmt"
	"log"

	"gin-starter/cmd/api/router"

	"gin-starter/cmd/api/config"
)

//	@Title			Gin Starter Swagger 2.0
//	@Version		1.0.0
//	@Description	Swagger API Documentation.

//	@BasePath	/api

func main() {
	if err := config.LoadConfig(".env"); err != nil {
		panic(fmt.Errorf("Missing env file: %s", err))
	}
	config.ConnectDB()

	r := router.RegisterRoutes()

	log.Fatalln(r.Run(fmt.Sprintf(":%v", config.Config.ServerPort)))
}
