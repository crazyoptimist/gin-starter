package auth

import (
	"fmt"
	"net/http"

	"gin-starter/internal/domain/model"
	"gin-starter/internal/domain/user"
	"gin-starter/pkg/utils"
)

type AuthHelper interface {
	BlacklistToken(token string) error
	IsTokenBlacklisted(token string) (bool, error)
}

var ErrTokenBlacklisted = fmt.Errorf("Refresh token is blacklisted")

type AuthService struct {
	UserRepository user.UserRepository
	AuthHelper     AuthHelper
}

func NewAuthService(userRepository user.UserRepository, authHelper AuthHelper) *AuthService {
	return &AuthService{
		UserRepository: userRepository,
		AuthHelper:     authHelper,
	}
}

func (s *AuthService) Register(registerDto *RegisterDto) (*LoginResponse, error) {
	if _, err := s.UserRepository.FindByEmail(registerDto.Email); err == nil {
		return nil, &utils.HttpError{
			Code:    http.StatusBadRequest,
			Message: "Account already exists with the email",
		}
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

	accessToken, refreshToken, err := GenerateTokenPair(user.ID)
	if err != nil {
		return nil, err
	}

	return &LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (s *AuthService) Login(loginDto *LoginDto) (*LoginResponse, error) {
	user, err := s.UserRepository.FindByEmail(loginDto.Email)
	if err != nil {
		return nil, err
	}

	if err := utils.VerifyPassword(user.Password, loginDto.Password); err != nil {
		return nil, &utils.HttpError{Code: http.StatusUnauthorized, Message: "Invalid password"}
	}

	accessToken, refreshToken, err := GenerateTokenPair(user.ID)
	if err != nil {
		return nil, err
	}

	return &LoginResponse{AccessToken: accessToken, RefreshToken: refreshToken}, nil
}

func (s *AuthService) Logout(logoutDto *LogoutDto) error {
	return s.AuthHelper.BlacklistToken(logoutDto.RefreshToken)
}

func (s *AuthService) RefreshToken(logoutDto *LogoutDto) (*LoginResponse, error) {

	isTokenValid, userId, keyId, err := ValidateToken(logoutDto.RefreshToken)
	if err != nil {
		return nil, err
	}

	if !isTokenValid {
		return nil, fmt.Errorf("Refresh token is invalid")
	}

	if keyId != RefreshTokenKeyId {
		return nil, fmt.Errorf("Wrong key ID")
	}

	isTokenBlacklisted, err := s.AuthHelper.IsTokenBlacklisted(logoutDto.RefreshToken)
	if err != nil {
		return nil, err
	}

	if isTokenBlacklisted {
		return nil, ErrTokenBlacklisted
	}

	err = s.AuthHelper.BlacklistToken(logoutDto.RefreshToken)
	if err != nil {
		return nil, err
	}

	accessToken, refreshToken, err := GenerateTokenPair(userId)
	if err != nil {
		return nil, err
	}

	return &LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}
