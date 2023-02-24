package auth

import (
	"gin-starter/internal/user"
)

type AuthService struct {
	UserRepository *user.UserRepository
}

func NewAuthService(repository *user.UserRepository) AuthService {
	return AuthService{UserRepository: repository}
}

func (s *AuthService) Register(user *user.User) (*LoginResponse, error) {
	user, err := s.UserRepository.Create(*user)
	if err != nil {
		return nil, err
	}

	accessToken, refreshToken, err := GenerateTokenPair(user.ID)
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
