package router

import (
	"github.com/gin-gonic/gin"

	"gin-starter/internal/config"
	"gin-starter/internal/controller"
)

func RegisterAuthRoutes(g *gin.RouterGroup) {
	controllers := controller.NewAuthController(config.Config.DB)
	g.POST("/register", controllers.Register)
	g.POST("/login", controllers.Login)
}

func RegisterUserRoutes(g *gin.RouterGroup) {
	controllers := controller.NewUserController(config.Config.DB)
	g.GET("", controllers.FindAll)
	g.GET("me", controllers.Me)
	g.GET(":id", controllers.FindById)
	g.POST("", controllers.Create)
	g.PATCH(":id", controllers.Update)
	g.DELETE(":id", controllers.Delete)
}
