package middleware

import (
	"strconv"

	"github.com/gin-gonic/gin"

	"gin-starter/internal/config"
)

func TotalCountMiddleware(modelToCount interface{}) gin.HandlerFunc {
	return func(c *gin.Context) {
		// TODO: Parse the query params for filters

		var count int64
		config.Config.DB.Model(modelToCount).Count(&count)

		c.Header("X-Total-Count", strconv.FormatInt(count, 10))

		c.Next()
	}
}
