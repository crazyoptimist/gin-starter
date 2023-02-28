package main

import (
	"fmt"
	"log"

	"gin-starter/internal/core/config"
	"gin-starter/internal/core/router"
)

// We are not using an API key here, but in OpenAPI v2 there is no better way to configure this.
// @securityDefinitions.apikey  JWT
// @in                          header
// @name                        Authorization

func main() {
	if err := config.LoadConfig(".env"); err != nil {
		panic(fmt.Errorf("Error while loading env file: %s", err))
	}
	config.ConnectDB()

	r := router.RegisterRoutes()
	router.SetupSwagger(r)

	log.Fatalln(r.Run(fmt.Sprintf(":%v", config.Config.ServerPort)))
}
