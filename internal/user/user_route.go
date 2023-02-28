package user

import (
	"github.com/gin-gonic/gin"

	"gin-starter/internal/core/config"
)

func RegisterRoutes(g *gin.RouterGroup) {
	controllers := NewUserController(config.Config.DB)
	g.GET("", controllers.FindAll)
	g.GET(":id", controllers.FindById)
	g.POST("", controllers.Create)
	g.PATCH(":id", controllers.Update)
	g.DELETE(":id", controllers.Delete)
}
