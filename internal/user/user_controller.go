package user

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"gin-starter/pkg/utils"
)

type UserController struct {
	UserService UserService
}

func NewUserController(db *gorm.DB) UserController {
	userRepository := NewUserRepository(db)
	userService := NewUserService(userRepository)
	return UserController{UserService: userService}
}

// FindAll godoc
// @Summary Retrieves users
// @Tags users
// @Success 200	{array} User
// @Failure 500 {object} utils.HttpError
// @Router /admin/users [get]
// @Security JWT
func (u *UserController) FindAll(c *gin.Context) {
	users := u.UserService.FindAll()
	c.JSON(http.StatusOK, users)
}

// FindById godoc
// @Summary Retrieves a user by ID
// @Tags users
// @Param id path integer true "User ID"
// @Success 200	{object} User
// @Failure 400 {object} utils.HttpError
// @Failure 404 {object} utils.HttpError
// @Failure 500 {object} utils.HttpError
// @Router /admin/users/{id} [get]
// @Security JWT
func (u *UserController) FindById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.RaiseHttpError(c, http.StatusBadRequest, err)
		return
	}

	user, err := u.UserService.FindById(uint(id))
	if err != nil {
		utils.RaiseHttpError(c, http.StatusNotFound, err)
		return
	}

	c.JSON(http.StatusOK, user)
}

// Create godoc
// @Summary Create a new user
// @Tags users
// @Param request body CreateUserDto true "CreateUserDto"
// @Success 201	{array} User
// @Failure 400 {object} utils.HttpError
// @Failure 500 {object} utils.HttpError
// @Router /admin/users [post]
// @Security JWT
func (u *UserController) Create(c *gin.Context) {
	var dto CreateUserDto
	if err := c.BindJSON(&dto); err != nil {
		utils.RaiseHttpError(c, http.StatusBadRequest, err)
		return
	}

	user, err := u.UserService.Create(&dto)
	if err != nil {
		utils.RaiseHttpError(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusCreated, user)
}

// Update godoc
// @Summary Update user
// @Tags users
// @Param id path integer true "User ID"
// @Param request body CreateUserDto true "UpdateUserDto"
// @Success 200	{array} User
// @Failure 400 {object} utils.HttpError
// @Failure 404 {object} utils.HttpError
// @Failure 500 {object} utils.HttpError
// @Router /admin/users/{id} [patch]
// @Security JWT
func (u *UserController) Update(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.RaiseHttpError(c, http.StatusBadRequest, err)
		return
	}

	if _, err := u.UserService.FindById(uint(id)); err != nil {
		utils.RaiseHttpError(c, http.StatusNotFound, err)
		return
	}

	var dto UpdateUserDto
	if err := c.BindJSON(&dto); err != nil {
		utils.RaiseHttpError(c, http.StatusBadRequest, err)
		return
	}

	user, err := u.UserService.Update(&dto, uint(id))
	if err != nil {
		utils.RaiseHttpError(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, user)
}

// Delete godoc
// @Summary Delete user
// @Tags users
// @Param id path integer true "User ID"
// @Success 200
// @Failure 500 {object} utils.HttpError
// @Router /admin/users/{id} [delete]
// @Security JWT
func (u *UserController) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.RaiseHttpError(c, http.StatusBadRequest, err)
		return
	}

	if err := u.UserService.Delete(uint(id)); err != nil {
		utils.RaiseHttpError(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, nil)
}
