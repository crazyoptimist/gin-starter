package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"gin-starter/internal/dto"
	"gin-starter/internal/model"
	"gin-starter/internal/repository"
	"gin-starter/internal/service"
	"gin-starter/pkg/utils"
)

type UserService interface {
	FindAll(queryParams utils.QueryParams) []model.User
	FindById(id uint) (*model.User, error)
	Create(createUserDto *dto.CreateUserDto) (*model.User, error)
	Update(updateUserDto *dto.UpdateUserDto, id uint) (*model.User, error)
	Delete(id uint) error
}

type userController struct {
	UserService
}

func NewUserController(db *gorm.DB) *userController {
	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	return &userController{UserService: userService}
}

// FindAll godoc
// @Summary Retrieves users
// @Tags users
// @Success 200	{array} model.User
// @Failure 500 {object} utils.HttpError
// @Router /users [get]
// @Security JWT
func (u *userController) FindAll(c *gin.Context) {
	queryParams := utils.GetQueryParams(c)
	users := u.UserService.FindAll(queryParams)
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
// @Router /users/{id} [get]
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
// @Success 201	{object} model.User
// @Failure 400 {object} utils.HttpError
// @Failure 500 {object} utils.HttpError
// @Router /users [post]
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

// Me godoc
// @Summary Get my profile
// @Tags auth
// @Success 200	{object} model.User
// @Failure 400 {object} utils.HttpError
// @Failure 500 {object} utils.HttpError
// @Router /users/me [post]
// @Security JWT
func (u *userController) Me(c *gin.Context) {
	id, _ := c.Get("user")

	user, err := u.UserService.FindById(uint(id.(int)))
	if err != nil {
		utils.RaiseHttpError(c, http.StatusNotFound, err)
		return
	}

	c.JSON(http.StatusOK, user)
}

// Update godoc
// @Summary Update user
// @Tags users
// @Param id path integer true "User ID"
// @Param request body dto.CreateUserDto true "UpdateUserDto"
// @Success 200	{object} model.User
// @Failure 400 {object} utils.HttpError
// @Failure 404 {object} utils.HttpError
// @Failure 500 {object} utils.HttpError
// @Router /users/{id} [patch]
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
// @Router /users/{id} [delete]
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
