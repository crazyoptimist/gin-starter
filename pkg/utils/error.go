package utils

import "github.com/gin-gonic/gin"

func RaiseHttpError(ctx *gin.Context, status int, err error) {
	ctx.JSON(status, HttpError{
		Code:    status,
		Message: err.Error(),
	})
}

type HttpError struct {
	Code    int    `json:"code" example:"400"`
	Message string `json:"message" example:"bad request"`
}

func (e *HttpError) Error() string {
	return e.Message
}
