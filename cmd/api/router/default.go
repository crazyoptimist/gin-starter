package router

import (
	"github.com/gin-gonic/gin"

	docs "gin-starter/docs"
	"gin-starter/internal/user"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func RegisterRoutes() *gin.Engine {
	router := gin.New()

	router.Use(gin.Logger())

	router.Use(gin.Recovery())

	api := router.Group("/api")
	{
		admin := api.Group("/admin")
		{
			users := admin.Group("/users")
			{
				user.RegisterRoutes(users)
			}
		}
	}

	return router
}

func SetupSwagger(r *gin.Engine) {
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	docs.SwaggerInfo.BasePath = "/api"
	docs.SwaggerInfo.Title = "Gin Starter Swagger 2.0"
	docs.SwaggerInfo.Version = "1.0.0"
	docs.SwaggerInfo.Description = "Swagger API Documentation"
}
