package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"gin-starter/internal/dto"
	"gin-starter/internal/repository"
	"gin-starter/internal/service"
	"gin-starter/pkg/utils"
)

type UserController interface {
	FindAll(c *gin.Context)
	FindById(c *gin.Context)
	Create(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}

type userController struct {
	UserService service.UserService
}

var _ UserController = (*userController)(nil)

func NewUserController(db *gorm.DB) UserController {
	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	return &userController{UserService: userService}
}

// FindAll godoc
// @Summary Retrieves users
// @Tags users
// @Success 200	{array} model.User
// @Failure 500 {object} utils.HttpError
// @Router /admin/users [get]
// @Security JWT
func (u *userController) FindAll(c *gin.Context) {
	users := u.UserService.FindAll()
	c.JSON(http.StatusOK, users)
}

// FindById godoc
// @Summary Retrieves a user by ID
// @Tags users
// @Param id path integer true "User ID"
// @Success 200	{object} model.User
// @Failure 400 {object} utils.HttpError
// @Failure 404 {object} utils.HttpError
// @Failure 500 {object} utils.HttpError
// @Router /admin/users/{id} [get]
// @Security JWT
func (u *userController) FindById(c *gin.Context) {
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
// @Param request body dto.CreateUserDto true "CreateUserDto"
// @Success 201	{array} model.User
// @Failure 400 {object} utils.HttpError
// @Failure 500 {object} utils.HttpError
// @Router /admin/users [post]
// @Security JWT
func (u *userController) Create(c *gin.Context) {
	var createUserDto dto.CreateUserDto
	if err := c.BindJSON(&createUserDto); err != nil {
		utils.RaiseHttpError(c, http.StatusBadRequest, err)
		return
	}

	user, err := u.UserService.Create(&createUserDto)
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
// @Param request body dto.CreateUserDto true "UpdateUserDto"
// @Success 200	{array} model.User
// @Failure 400 {object} utils.HttpError
// @Failure 404 {object} utils.HttpError
// @Failure 500 {object} utils.HttpError
// @Router /admin/users/{id} [patch]
// @Security JWT
func (u *userController) Update(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.RaiseHttpError(c, http.StatusBadRequest, err)
		return
	}

	if _, err := u.UserService.FindById(uint(id)); err != nil {
		utils.RaiseHttpError(c, http.StatusNotFound, err)
		return
	}

	var updateUserDto dto.UpdateUserDto
	if err := c.BindJSON(&updateUserDto); err != nil {
		utils.RaiseHttpError(c, http.StatusBadRequest, err)
		return
	}

	user, err := u.UserService.Update(&updateUserDto, uint(id))
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
func (u *userController) Delete(c *gin.Context) {
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
