package user

import (
	"gin-starter/pkg/utils"
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

// Create godoc
// @Summary Create a new user
// @Tags users
// @Param request body CreateUserDto true "CreateUserDto"
// @Success 200	{array} User
// @Failure 400 {object} utils.HTTPError
// @Failure 500 {object} utils.HTTPError
// @Router /admin/users [post]
func (u *UserController) Create(c *gin.Context) {
	var dto CreateUserDto
	if err := c.BindJSON(&dto); err != nil {
		utils.NewError(c, http.StatusBadRequest, err)
	}

	user, err := u.UserService.Save(MapCreateUserDto(dto))
	if err != nil {
		utils.NewError(c, http.StatusInternalServerError, err)
	}

	c.JSON(http.StatusCreated, user)
}

// FindAll godoc
// @Summary Retrieves users
// @Tags users
// @Success 200	{array} User
// @Failure 500 {object} utils.HTTPError
// @Router /admin/users [get]
func (u *UserController) FindAll(c *gin.Context) {
	users := u.UserService.FindAll()
	c.JSON(http.StatusOK, users)
}

// FindById godoc
// @Summary Retrieves a user by ID
// @Tags users
// @Param id path integer true "User ID"
// @Success 200	{object} User
// @Failure 400 {object} utils.HTTPError
// @Failure 404 {object} utils.HTTPError
// @Failure 500 {object} utils.HTTPError
// @Router /admin/users/{id} [get]
func (u *UserController) FindById(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.NewError(c, http.StatusBadRequest, err)
	}

	user, err := u.UserService.FindById(uint(id))
	if err != nil {
		utils.NewError(c, http.StatusNotFound, err)
		return
	}

	c.JSON(http.StatusOK, user)
}
