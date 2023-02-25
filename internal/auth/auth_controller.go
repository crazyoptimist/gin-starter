package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"gin-starter/internal/user"
	"gin-starter/pkg/utils"
)

type AuthController struct {
	AuthService AuthService
}

func NewAuthController(db *gorm.DB) AuthController {
	userRepository := user.NewUserRepository(db)
	authService := NewAuthService(userRepository)
	return AuthController{AuthService: authService}
}

// Register godoc
// @Summary Register new user
// @Tags auth
// @Param request body RegisterDto true "RegisterDto"
// @Success 201	{array} LoginResponse
// @Failure 400 {object} utils.HttpError
// @Failure 500 {object} utils.HttpError
// @Router /auth/register [post]
func (a *AuthController) Register(c *gin.Context) {
	var dto RegisterDto
	if err := c.BindJSON(&dto); err != nil {
		utils.RaiseHttpError(c, http.StatusBadRequest, err)
		return
	}

	mappedUser, err := MapRegisterDto(&dto)
	if err != nil {
		utils.RaiseHttpError(c, http.StatusBadRequest, err)
		return
	}

	loginResponse, err := a.AuthService.Register(&mappedUser)
	if err != nil {
		utils.RaiseHttpError(c, http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusCreated, loginResponse)
}
