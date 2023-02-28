package auth

import (
	"net/http"

	"gin-starter/internal/user"
	"gin-starter/pkg/utils"
)

type AuthService struct {
	UserRepository *user.UserRepository
}

func NewAuthService(repository *user.UserRepository) AuthService {
	return AuthService{UserRepository: repository}
}

func (s *AuthService) Register(dto *RegisterDto) (*LoginResponse, error) {
	if _, err := s.UserRepository.FindByEmail(dto.Email); err == nil {
		return nil, &utils.HttpError{Code: http.StatusBadRequest, Message: "Account already exists with the email"}
	}

	if hashedPassword, err := utils.HashPassword(dto.Password); err != nil {
		return nil, err
	} else {
		dto.Password = hashedPassword
	}

	mappedUser, err := MapRegisterDto(dto)
	if err != nil {
		return nil, err
	}

	user, err := s.UserRepository.Create(mappedUser)
	if err != nil {
		return nil, err
	}

	accessToken, refreshToken, err := GenerateTokenPair(user.ID)
	if err != nil {
		return nil, err
	}

	return &LoginResponse{accessToken, refreshToken}, nil
}

func (s *AuthService) Login(dto *LoginDto) (*LoginResponse, error) {
	user, err := s.UserRepository.FindByEmail(dto.Email)
	if err != nil {
		return nil, err
	}

	if err := utils.VerifyPassword(user.Password, dto.Password); err != nil {
		return nil, &utils.HttpError{Code: http.StatusUnauthorized, Message: "Invalid password"}
	}

	accessToken, refreshToken, err := GenerateTokenPair(user.ID)
	if err != nil {
		return nil, err
	}

	return &LoginResponse{accessToken, refreshToken}, nil
}

func (s *AuthService) Logout() {
	// TODO: Implement it
	return
}
