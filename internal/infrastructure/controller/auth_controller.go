package controller

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"gin-starter/internal/domain/auth"
	"gin-starter/internal/domain/model"
	"gin-starter/internal/infrastructure/repository"
	"gin-starter/pkg/common"
)

type authController struct {
	AuthService auth.AuthService
}

func NewAuthController(db *gorm.DB) *authController {
	userRepository := repository.NewUserRepository(db)
	authService := auth.NewAuthService(userRepository)
	return &authController{AuthService: *authService}
}

// Register godoc
// @Summary Register new user
// @Tags auth
// @Param request body auth.RegisterDto true "RegisterDto"
// @Success 201	{object} auth.LoginResponse
// @Failure 400 {object} common.HttpError
// @Failure 500 {object} common.HttpError
// @Router /auth/register [post]
func (a *authController) Register(c *gin.Context) {
	var registerDto auth.RegisterDto
	if err := c.BindJSON(&registerDto); err != nil {
		common.RaiseHttpError(c, http.StatusBadRequest, err)
		return
	}

	loginResponse, err := a.AuthService.Register(&registerDto)
	if err != nil {
		common.RaiseHttpError(c, http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusCreated, loginResponse)
}

// Login godoc
// @Summary Login user
// @Tags auth
// @Param request body auth.LoginDto true "LoginDto"
// @Success 201	{object} auth.LoginResponse
// @Failure 400 {object} common.HttpError
// @Failure 401 {object} common.HttpError
// @Failure 404 {object} common.HttpError
// @Router /auth/login [post]
func (a *authController) Login(c *gin.Context) {
	var loginDto auth.LoginDto
	if err := c.BindJSON(&loginDto); err != nil {
		common.RaiseHttpError(c, http.StatusBadRequest, err)
		return
	}

	loginResponse, err := a.AuthService.Login(&loginDto)
	if err != nil {
		common.RaiseHttpError(c, http.StatusUnauthorized, err)
		return
	}

	c.JSON(http.StatusCreated, loginResponse)
}

// Logout godoc
// @Summary Logout user
// @Tags auth
// @Success 200
// @Failure 401 {object} common.HttpError
// @Failure 500 {object} common.HttpError
// @Router /auth/logout [post]
// @Security JWT
func (a *authController) Logout(c *gin.Context) {
	user, exists := c.Get("user")
	if !exists {
		common.RaiseHttpError(
			c,
			http.StatusInternalServerError,
			errors.New("User missing in context"),
		)
		return
	}

	if err := a.AuthService.Logout(user.(model.User)); err != nil {
		common.RaiseHttpError(c, http.StatusInternalServerError, err)
		return
	}

	c.Status(http.StatusOK)
}

// RefreshToken godoc
// @Summary Refresh tokens and invalidate the previous refresh token
// @Tags auth
// @Param request body auth.LogoutDto true "TokenRefresh DTO"
// @Success 201	{object} auth.LoginResponse
// @Failure 400 {object} common.HttpError
// @Failure 500 {object} common.HttpError
// @Router /auth/refresh [post]
func (a *authController) Refresh(c *gin.Context) {
	var refreshDto auth.RefreshDto
	if err := c.BindJSON(&refreshDto); err != nil {
		common.RaiseHttpError(c, http.StatusBadRequest, err)
		return
	}

	refreshResponse, err := a.AuthService.Refresh(&refreshDto)
	if err != nil {
		common.RaiseHttpError(c, http.StatusUnauthorized, err)
		return
	}

	c.JSON(http.StatusCreated, refreshResponse)
}
