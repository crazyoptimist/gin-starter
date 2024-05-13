package utils

import (
	"slices"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type PaginationParam struct {
	Offset int
	Limit  int
}

func GetPaginationParam(c *gin.Context) PaginationParam {
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

	return PaginationParam{
		Offset: offset,
		Limit:  limit,
	}
}

type SortParam struct {
	FieldName string
	Order     string
}

func GetSortParams(c *gin.Context) []SortParam {
	fieldsString, hasSortFields := c.GetQuery("_sort")
	ordersString, hasSortOrders := c.GetQuery("_order")

	if !hasSortFields || !hasSortOrders {
		return []SortParam{}
	}

	fields := strings.Split(fieldsString, ",")
	orders := strings.Split(ordersString, ",")

	// We need to make sure the length of fields and orders are
	// the same or take the smaller length
	fieldsLen := len(fields)
	ordersLen := len(orders)

	sortersLength := fieldsLen

	if fieldsLen > ordersLen {
		sortersLength = ordersLen
	}

	var sortParams = []SortParam{}

	for i := 0; i < sortersLength; i++ {
		order := strings.ToUpper(orders[i])

		if !(order == "ASC" || order == "DESC") {
			continue
		}

		sortParams = append(sortParams, SortParam{
			FieldName: fields[i],
			Order:     orders[i],
		})
	}

	return sortParams
}

// We must handle the actual value types in the query composition
// We only support "eq" operator for now (add a field to support different operators)
type FilterParam struct {
	FieldName string
	Value     string
}

func GetFilterParams(c *gin.Context) []FilterParam {

	var filterParams = []FilterParam{}

	queryMap := c.Request.URL.Query()
	paginationAndSortKeys := []string{"_limit", "_offset", "_sort", "_order"}

	for key, val := range queryMap {
		if slices.Contains(paginationAndSortKeys, key) {
			continue
		}
		filterParams = append(filterParams, FilterParam{
			FieldName: key,
			Value:     val[0],
		})
	}

	return filterParams
}
