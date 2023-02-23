package user

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserController struct {
	UserService UserService
}

func NewUserController(db *gorm.DB) UserController {
	userRepository := NewUserRepository(db)
	userService := NewUserService(userRepository)
	return UserController{UserService: userService}
}

// GetUser godoc
// @Summary Retrieves a user by ID
// @Tags users
// @Param id path integer true "User ID"
// @Success 200	{object} User
// @Router /users/{id} [get]
func (u *UserController) FindById(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	if user, err := u.UserService.FindById(uint(id)); err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		log.Println(err)
	} else {
		c.JSON(http.StatusOK, user)
	}
}
