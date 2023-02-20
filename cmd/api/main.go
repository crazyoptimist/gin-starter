package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"

	docs "gin-starter/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"gin-starter/cmd/api/config"
	"gin-starter/internal/user"
)

// @title Gin Starter Swagger 2.0
// @version 1.0.0
// @description Swagger API Documentation.

// @BasePath /api/v1
func main() {
	if err := config.LoadConfig("./configs"); err != nil {
		panic(fmt.Errorf("invalid application configuration: %s", err))
	}
	config.ConnectDB()

	r := gin.New()

	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	docs.SwaggerInfo.BasePath = "/api/v1"

	v1 := r.Group("/api/v1")
	{
		users := v1.Group("/users")
		{
			users.GET(":id", user.GetUser)
		}
	}

	log.Fatalln(r.Run(fmt.Sprintf(":%v", config.Config.ServerPort)))
}
