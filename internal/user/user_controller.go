package user

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetUser godoc
// @Summary Retrieves a user by ID
// @Tags users
// @Produce json
// @Param id path integer true "User ID"
// @Success 200	{object} User
// @Router /users/{id} [get]
func GetUser(c *gin.Context) {
	s := NewUserService(NewUserDAO())
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	if user, err := s.Get(uint(id)); err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		log.Println(err)
	} else {
		c.JSON(http.StatusOK, user)
	}
}
