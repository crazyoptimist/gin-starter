package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"gin-starter/internal/domain/user"
	"gin-starter/internal/infrastructure/repository"
	"gin-starter/pkg/common"
	"gin-starter/pkg/utils"
)

type userController struct {
	UserService user.UserService
}

func NewUserController(db *gorm.DB) *userController {
	userRepository := repository.NewUserRepository(db)
	userService := user.NewUserService(userRepository)
	return &userController{UserService: *userService}
}

// FindAll godoc
// @Summary Retrieves users
// @Tags users
// @Success 200	{array} model.User
// @Failure 500 {object} utils.HttpError
// @Router /users [get]
// @Security JWT
func (u *userController) FindAll(c *gin.Context) {
	paginationParam := common.GetPaginationParam(c)
	sortParams := common.GetSortParams(c)
	filterParams := common.GetFilterParams(c)

	users, totalCount, err := u.UserService.FindAll(
		paginationParam,
		sortParams,
		filterParams,
	)
	if err != nil {
		utils.RaiseHttpError(c, http.StatusInternalServerError, err)
		return
	}

	c.Header("X-Total-Count", strconv.FormatInt(totalCount, 10))

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
// @Param request body user.CreateUserDto true "CreateUserDto"
// @Success 201	{object} model.User
// @Failure 400 {object} utils.HttpError
// @Failure 500 {object} utils.HttpError
// @Router /users [post]
// @Security JWT
func (u *userController) Create(c *gin.Context) {
	var createUserDto user.CreateUserDto
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
// @Tags users
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
// @Param request body user.UpdateUserDto true "UpdateUserDto"
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

	var updateUserDto user.UpdateUserDto
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
