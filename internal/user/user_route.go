package user

import (
	"gin-starter/cmd/api/config"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(g *gin.RouterGroup) {
	controllers := NewUserController(config.Config.DB)
	g.GET(":id", controllers.FindById)
}
