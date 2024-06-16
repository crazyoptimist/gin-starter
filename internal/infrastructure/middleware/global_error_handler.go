package middleware

import (
	"errors"
	"net/http"

	"gin-starter/pkg/common"

	"github.com/gin-gonic/gin"
)

func GlobalErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		for _, err := range c.Errors {
			switch e := err.Err.(type) {
			case *common.HttpError:
				c.AbortWithStatusJSON(e.StatusCode, e)
			default:
				common.RaiseHttpError(
					c,
					http.StatusInternalServerError,
					errors.New("Service Unavailable"),
				)
			}
		}
	}
}
