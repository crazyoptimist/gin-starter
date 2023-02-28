package auth

import (
	"net/http"

	"gin-starter/internal/user"
	"gin-starter/pkg/utils"
)

type LoginResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type LoginDto struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type RegisterDto struct {
	FirstName            string `json:"firstName"`
	LastName             string `json:"lastName"`
	Email                string `json:"email" binding:"required"`
	Password             string `json:"password" binding:"required"`
	PasswordConfirmation string `json:"passwordConfirmation" binding:"required"`
}

func MapRegisterDto(dto *RegisterDto) (user.User, error) {
	if dto.Password != dto.PasswordConfirmation {
		return user.User{}, &utils.HttpError{Code: http.StatusBadRequest, Message: "Password mismatched"}
	}

	return user.User{
		FirstName: dto.FirstName,
		LastName:  dto.LastName,
		Email:     dto.Email,
		Password:  dto.Password,
	}, nil
}
