package router

import (
	"github.com/gin-gonic/gin"

	docs "gin-starter/docs"
	"gin-starter/internal/auth"
	"gin-starter/internal/user"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func RegisterRoutes() *gin.Engine {
	router := gin.New()

	router.Use(gin.Logger())

	router.Use(gin.Recovery())

	v1 := router.Group("/api")
	{
		authGroup := v1.Group("/auth")
		{
			auth.RegisterRoutes(authGroup)
		}
		adminGroup := v1.Group("/admin", auth.AuthMiddleware())
		{
			usersGroup := adminGroup.Group("/users")
			{
				user.RegisterRoutes(usersGroup)
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
