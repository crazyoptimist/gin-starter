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

func RegisterRoutes() *gin.Engine {
	if err := config.LoadConfig(".env"); err != nil {
		panic(fmt.Errorf("Missing env file: %s", err))
	}
	config.ConnectDB()

	r := gin.New()

	r.Use(gin.Logger())

	r.Use(gin.Recovery())

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

func SetupSwagger(r *gin.Engine) {
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	docs.SwaggerInfo.BasePath = "/api"
	docs.SwaggerInfo.Title = "Gin Starter Swagger 2.0"
	docs.SwaggerInfo.Version = "1.0.0"
	docs.SwaggerInfo.Description = "Swagger API Documentation"
}
