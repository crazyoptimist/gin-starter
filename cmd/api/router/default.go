package router

import (
	"fmt"

	"github.com/gin-gonic/gin"

	docs "gin-starter/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"gin-starter/cmd/api/config"
	"gin-starter/internal/user"
)

//	@Title			Gin Starter Swagger 2.0
//	@Version		1.0.0
//	@Description	Swagger API Documentation.

//	@BasePath	/api

func RegisterRoutes() *gin.Engine {
	if err := config.LoadConfig(".env"); err != nil {
		panic(fmt.Errorf("Missing env file: %s", err))
	}
	config.ConnectDB()

	r := gin.New()

	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	docs.SwaggerInfo.BasePath = "/api"

	api := r.Group("/api")
	{
		admin := api.Group("/admin")
		{
			users := admin.Group("/users")
			{
				users.GET(":id", user.GetUser)
			}
		}
	}

	return r
}
