package router

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	docs "gin-starter/docs"
	"gin-starter/internal/infrastructure/middleware"
)

func RegisterRoutes() *gin.Engine {
	router := gin.New()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	corsConfig.AllowHeaders = []string{
		"Authorization",
		"Content-Type",
		"Access-Control-Allow-Origin",
		"X-CSRF-Token",
	}
	corsConfig.ExposeHeaders = []string{
		"X-Total-Count",
	}
	router.Use(cors.New(corsConfig))

	// Health check route
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "OK")
	})

	v1 := router.Group("/api")
	{
		authGroup := v1.Group("/auth")
		RegisterAuthRoutes(authGroup)
		usersGroup := v1.Group("/users", middleware.AuthMiddleware())
		RegisterUserRoutes(usersGroup)
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
