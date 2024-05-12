package controller

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"gin-starter/internal/dto"
	"gin-starter/internal/repository"
	"gin-starter/internal/service"
	"gin-starter/pkg/utils"
)

type AuthService interface {
	Register(registerDto *dto.RegisterDto) (*dto.LoginResponse, error)
	Login(loginDto *dto.LoginDto) (*dto.LoginResponse, error)
	Logout(logoutDto *dto.LogoutDto) error
	RefreshToken(logoutDto *dto.LogoutDto) (*dto.LoginResponse, error)
}

type authController struct {
	AuthService
}

func NewAuthController(db *gorm.DB) *authController {
	userRepository := repository.NewUserRepository(db)
	authService := service.NewAuthService(userRepository)
	return &authController{AuthService: authService}
}

// Register godoc
// @Summary Register new user
// @Tags auth
// @Param request body dto.RegisterDto true "RegisterDto"
// @Success 201	{object} dto.LoginResponse
// @Failure 400 {object} utils.HttpError
// @Failure 500 {object} utils.HttpError
// @Router /auth/register [post]
func (a *authController) Register(c *gin.Context) {
	var registerDto dto.RegisterDto
	if err := c.BindJSON(&registerDto); err != nil {
		utils.RaiseHttpError(c, http.StatusBadRequest, err)
		return
	}

	loginResponse, err := a.AuthService.Register(&registerDto)
	if err != nil {
		utils.RaiseHttpError(c, http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusCreated, loginResponse)
}

// Login godoc
// @Summary Login user
// @Tags auth
// @Param request body dto.LoginDto true "LoginDto"
// @Success 201	{object} dto.LoginResponse
// @Failure 400 {object} utils.HttpError
// @Failure 401 {object} utils.HttpError
// @Failure 404 {object} utils.HttpError
// @Router /auth/login [post]
func (a *authController) Login(c *gin.Context) {
	var loginDto dto.LoginDto
	if err := c.BindJSON(&loginDto); err != nil {
		utils.RaiseHttpError(c, http.StatusBadRequest, err)
		return
	}

	loginResponse, err := a.AuthService.Login(&loginDto)
	if err != nil {
		utils.RaiseHttpError(c, http.StatusUnauthorized, err)
		return
	}

	c.JSON(http.StatusCreated, loginResponse)
}

// Logout godoc
// @Summary Logout user (Invalidates refresh token)
// @Tags auth
// @Success 200
// @Failure 401 {object} utils.HttpError
// @Failure 500 {object} utils.HttpError
// @Router /auth/logout [post]
// @Security JWT
func (a *authController) Logout(c *gin.Context) {
	var logoutDto dto.LogoutDto
	if err := c.BindJSON(&logoutDto); err != nil {
		utils.RaiseHttpError(c, http.StatusBadRequest, err)
		return
	}

	if err := a.AuthService.Logout(&logoutDto); err != nil {
		utils.RaiseHttpError(c, http.StatusInternalServerError, err)
		return
	}

	c.Status(http.StatusOK)
}

// RefreshToken godoc
// @Summary Refresh tokens and invalidate the previous refresh token
// @Tags auth
// @Param request body dto.Logout true "TokenRefresh DTO"
// @Success 201	{object} dto.LoginResponse
// @Failure 400 {object} utils.HttpError
// @Failure 500 {object} utils.HttpError
// @Router /auth/refresh [post]
func (a *authController) RefreshToken(c *gin.Context) {
	var refreshDto dto.LogoutDto
	if err := c.BindJSON(&refreshDto); err != nil {
		utils.RaiseHttpError(c, http.StatusBadRequest, err)
		return
	}

	refreshResponse, err := a.AuthService.RefreshToken(&refreshDto)
	if err != nil {
		if errors.Is(err, service.ErrTokenBlacklisted) {
			utils.RaiseHttpError(c, http.StatusUnauthorized, err)
			return
		}
		utils.RaiseHttpError(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusCreated, refreshResponse)
}
