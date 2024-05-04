package utils

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

type QueryParams struct {
	Offset int
	Limit  int
}

func GetQueryParams(c *gin.Context) QueryParams {
	offsetString := c.DefaultQuery("_offset", "0")
	limitString := c.DefaultQuery("_limit", "25")

	offset, err := strconv.Atoi(offsetString)
	if err != nil {
		offset = 0
	}
	limit, err := strconv.Atoi(limitString)
	if err != nil {
		limit = 25
	}

	return QueryParams{
		Offset: offset,
		Limit:  limit,
	}
}
