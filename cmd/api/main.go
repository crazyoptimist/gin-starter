package main

import (
	"fmt"
	"log"

	"gin-starter/cmd/api/config"
	"gin-starter/internal/user"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	docs "gin-starter/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Gin Starter Swagger 2.0
// @version 1.0.0
// @description Swagger API Documentation.

// @BasePath /api/v1
func main() {
	if err := config.LoadConfig("./configs"); err != nil {
		panic(fmt.Errorf("invalid application configuration: %s", err))
	}

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

	var dbErr error
	config.Config.DB, dbErr = gorm.Open("postgres", config.Config.DSN)
	if dbErr != nil {
		panic(dbErr)
	}

	// config.Config.DB.AutoMigrate(&models.User{})

	defer config.Config.DB.Close()

	log.Println("Successfully connected to database")

	r.Run(fmt.Sprintf(":%v", config.Config.ServerPort))
}
