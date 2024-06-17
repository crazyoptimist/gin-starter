package auth

import (
	"errors"
	"strconv"

	"gin-starter/internal/config"
	"gin-starter/internal/domain/model"
	"gin-starter/pkg/common"
	"gin-starter/pkg/utils"
)

type UserRepository interface {
	FindAll(
		paginationParam common.PaginationParam,
		sortParams []common.SortParam,
		filterParams []common.FilterParam,
	) ([]model.User, int64, error)
	FindById(id int) (*model.User, error)
	FindByEmail(email string) (*model.User, error)
	Create(user model.User) (*model.User, error)
	Update(user model.User) (*model.User, error)
	Delete(user model.User) error
	UpdateRefreshToken(userId int, refreshToken string) error
}

type AuthService struct {
	UserRepository UserRepository
}

func NewAuthService(userRepository UserRepository) *AuthService {
	return &AuthService{
		UserRepository: userRepository,
	}
}

func (s *AuthService) Register(registerDto *RegisterDto) (*LoginResponse, error) {
	if _, err := s.UserRepository.FindByEmail(registerDto.Email); err == nil {
		return nil, errors.New("Account already exists with the same email")
	}

	if registerDto.Password != registerDto.PasswordConfirmation {
		return nil, errors.New("Password mismatched")
	}

	if hashedPassword, err := utils.HashPassword(registerDto.Password); err != nil {
		return nil, err
	} else {
		registerDto.Password = hashedPassword
	}

	user, err := s.UserRepository.Create(model.User{
		FirstName: &registerDto.FirstName,
		LastName:  &registerDto.LastName,
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
		ExpiresIn:    config.Global.JwtAccessTokenExpiresIn.Seconds(),
	}, nil
}

func (s *AuthService) Login(loginDto *LoginDto) (*LoginResponse, error) {
	user, err := s.UserRepository.FindByEmail(loginDto.Email)
	if err != nil {
		return nil, err
	}

	if err := utils.VerifyPassword(user.Password, loginDto.Password); err != nil {
		return nil, errors.New("Invalid password")
	}

	accessToken, refreshToken, err := GenerateTokenPair(user.ID)
	if err != nil {
		return nil, err
	}

	if err := s.UserRepository.UpdateRefreshToken(
		user.ID,
		refreshToken,
	); err != nil {
		return nil, err
	}

	return &LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		ExpiresIn:    config.Global.JwtAccessTokenExpiresIn.Seconds(),
	}, nil
}

func (s *AuthService) Logout(user model.User) error {
	err := s.UserRepository.UpdateRefreshToken(user.ID, "")
	return err
}

func (s *AuthService) Refresh(dto *RefreshDto) (*LoginResponse, error) {

	_, userIdString, _, err := ValidateJwtToken(dto.RefreshToken)
	if err != nil {
		return nil, err
	}

	userId, err := strconv.Atoi(userIdString)
	if err != nil {
		return nil, err
	}

	user, err := s.UserRepository.FindById(userId)
	if err != nil {
		return nil, err
	}

	if *user.RefreshToken != dto.RefreshToken {
		return nil, errors.New("Invalid refresh token")
	}

	accessToken, refreshToken, err := GenerateTokenPair(userId)
	if err != nil {
		return nil, err
	}

	if err := s.UserRepository.UpdateRefreshToken(
		userId,
		refreshToken,
	); err != nil {
		return nil, err
	}

	return &LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		ExpiresIn:    config.Global.JwtAccessTokenExpiresIn.Seconds(),
	}, nil
}
