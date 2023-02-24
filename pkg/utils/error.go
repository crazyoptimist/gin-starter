package utils

import "github.com/gin-gonic/gin"

func RaiseHttpError(ctx *gin.Context, status int, err error) {
	ctx.JSON(status, ErrorResponse{
		Code:    status,
		Message: err.Error(),
	})
}

type ErrorResponse struct {
	Code    int    `json:"code" example:"400"`
	Message string `json:"message" example:"bad request"`
}
