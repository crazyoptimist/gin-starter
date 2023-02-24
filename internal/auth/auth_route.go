package auth

import (
	"gin-starter/cmd/api/config"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(g *gin.RouterGroup) {
	controllers := NewAuthController(config.Config.DB)
	g.POST("/register", controllers.Register)
}
