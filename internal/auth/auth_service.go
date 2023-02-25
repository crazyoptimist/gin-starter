package auth

import (
	"gin-starter/internal/user"
	"gin-starter/pkg/utils"
	"net/http"
)

type AuthService struct {
	UserRepository *user.UserRepository
}

func NewAuthService(repository *user.UserRepository) AuthService {
	return AuthService{UserRepository: repository}
}

func (s *AuthService) Register(user *user.User) (*LoginResponse, error) {
	if _, err := s.UserRepository.FindByEmail(user.Email); err == nil {
		return nil, &utils.HttpError{Code: http.StatusBadRequest, Message: "Account already exists with the email"}
	}

	if hashedPassword, err := utils.HashPassword(user.Password); err != nil {
		return nil, err
	} else {
		user.Password = hashedPassword
	}

	newUser, err := s.UserRepository.Create(*user)
	if err != nil {
		return nil, err
	}

	accessToken, refreshToken, err := GenerateTokenPair(newUser.ID)
	if err != nil {
		return nil, err
	}

	return &LoginResponse{accessToken, refreshToken}, nil
}

func (s *AuthService) Login() (*LoginResponse, error) {
	return &LoginResponse{}, nil
}

func (s *AuthService) Logout() {
	return
}
