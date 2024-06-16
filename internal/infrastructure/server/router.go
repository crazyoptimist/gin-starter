package server

import (
	"github.com/gin-gonic/gin"

	"gin-starter/internal/config"
	"gin-starter/internal/infrastructure/controller"
	"gin-starter/internal/infrastructure/middleware"
)

func registerAuthRoutes(g *gin.RouterGroup) {
	controllers := controller.NewAuthController(config.Global.DB)
	g.POST("/register", controllers.Register)
	g.POST("/login", controllers.Login)
	g.POST("/logout", controllers.Logout, middleware.AuthMiddleware())
	g.POST("/refresh", controllers.Refresh)
}

func registerUserRoutes(g *gin.RouterGroup) {
	controllers := controller.NewUserController(config.Global.DB)
	g.GET("", controllers.FindAll)
	g.GET("me", controllers.Me)
	g.GET(":id", controllers.FindById)
	g.POST("", controllers.Create)
	g.PATCH(":id", controllers.Update)
	g.DELETE(":id", controllers.Delete)
}
