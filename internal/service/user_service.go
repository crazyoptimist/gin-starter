package service

import (
	"gin-starter/internal/dto"
	"gin-starter/internal/model"
	"gin-starter/internal/repository"
	"gin-starter/pkg/common"
)

type UserService interface {
	FindAll() []model.User
	FindById(id uint) (*model.User, error)
	Create(createUserDto *dto.CreateUserDto) (*model.User, error)
	Update(updateUserDto *dto.UpdateUserDto, id uint) (*model.User, error)
	Delete(id uint) error
}

type userService struct {
	UserRepository repository.UserRepository
}

var _ UserService = (*userService)(nil)

func NewUserService(userRepository repository.UserRepository) UserService {
	return &userService{UserRepository: userRepository}
}

func (u *userService) FindAll() []model.User {
	return u.UserRepository.FindAll()
}

func (u *userService) FindById(id uint) (*model.User, error) {
	return u.UserRepository.FindById(id)
}

func (u *userService) Create(createUserDto *dto.CreateUserDto) (*model.User, error) {
	return u.UserRepository.Create(dto.MapCreateUserDto(createUserDto))
}

func (u *userService) Update(updateUserDto *dto.UpdateUserDto, id uint) (*model.User, error) {
	return u.UserRepository.Update(dto.MapUpdateUserDto(updateUserDto, id))
}

func (u *userService) Delete(id uint) error {
	return u.UserRepository.Delete(model.User{BaseModel: common.BaseModel{ID: id}})
}
