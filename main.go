package main

import (
	"fmt"
	"log"

	"gin-starter/cmd/api/config"
	"gin-starter/cmd/api/router"
)

func main() {
	if err := config.LoadConfig(".env"); err != nil {
		panic(fmt.Errorf("Missing env file: %s", err))
	}
	config.ConnectDB()

	r := router.RegisterRoutes()
	router.SetupSwagger(r)

	log.Fatalln(r.Run(fmt.Sprintf(":%v", config.Config.ServerPort)))
}
