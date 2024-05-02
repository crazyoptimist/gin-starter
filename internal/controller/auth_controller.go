package controller

import (
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
	Logout()
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
