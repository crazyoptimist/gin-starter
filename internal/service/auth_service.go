package service

import (
	"net/http"

	"gin-starter/internal/dto"
	"gin-starter/internal/helper"
	"gin-starter/internal/model"
	"gin-starter/pkg/utils"
)

type authService struct {
	UserRepository
}

func NewAuthService(userRepository UserRepository) *authService {
	return &authService{UserRepository: userRepository}
}

func (s *authService) Register(registerDto *dto.RegisterDto) (*dto.LoginResponse, error) {
	if _, err := s.UserRepository.FindByEmail(registerDto.Email); err == nil {
		return nil, &utils.HttpError{Code: http.StatusBadRequest, Message: "Account already exists with the email"}
	}

	if registerDto.Password != registerDto.PasswordConfirmation {
		return nil, &utils.HttpError{Code: http.StatusBadRequest, Message: "Password mismatched"}
	}

	if hashedPassword, err := utils.HashPassword(registerDto.Password); err != nil {
		return nil, err
	} else {
		registerDto.Password = hashedPassword
	}

	user, err := s.UserRepository.Create(model.User{
		FirstName: registerDto.FirstName,
		LastName:  registerDto.LastName,
		Email:     registerDto.Email,
		Password:  registerDto.Password,
	})
	if err != nil {
		return nil, err
	}

	accessToken, refreshToken, err := helper.GenerateTokenPair(user.ID)
	if err != nil {
		return nil, err
	}

	return &dto.LoginResponse{AccessToken: accessToken, RefreshToken: refreshToken}, nil
}

func (s *authService) Login(loginDto *dto.LoginDto) (*dto.LoginResponse, error) {
	user, err := s.UserRepository.FindByEmail(loginDto.Email)
	if err != nil {
		return nil, err
	}

	if err := utils.VerifyPassword(user.Password, loginDto.Password); err != nil {
		return nil, &utils.HttpError{Code: http.StatusUnauthorized, Message: "Invalid password"}
	}

	accessToken, refreshToken, err := helper.GenerateTokenPair(user.ID)
	if err != nil {
		return nil, err
	}

	return &dto.LoginResponse{AccessToken: accessToken, RefreshToken: refreshToken}, nil
}

func (s *authService) Logout() {
	// TODO: Implement this
	return
}
