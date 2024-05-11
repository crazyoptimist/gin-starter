package service

import (
	"gin-starter/internal/dto"
	"gin-starter/internal/model"
	"gin-starter/pkg/common"
	"gin-starter/pkg/utils"
)

// IMPORTANT: Always define an interface where it is used (injected)!
// DO NOT define it where it is implemented!
type UserRepository interface {
	FindAll(
		paginationParam utils.PaginationParam,
		sortParams []utils.SortParam,
		filterParams []utils.FilterParam,
	) ([]model.User, int64, error)
	FindById(id uint) (*model.User, error)
	FindByEmail(email string) (*model.User, error)
	Create(user model.User) (*model.User, error)
	Update(user model.User) (*model.User, error)
	Delete(user model.User) error
}

// Why don't inject the repository into the controller?
// Adding business logic to the service is better.
type userService struct {
	UserRepository UserRepository
}

func NewUserService(userRepository UserRepository) *userService {
	return &userService{UserRepository: userRepository}
}

func (u *userService) FindAll(
	paginationParam utils.PaginationParam,
	sortParams []utils.SortParam,
	filterParams []utils.FilterParam,
) ([]model.User, int64, error) {
	return u.UserRepository.FindAll(paginationParam, sortParams, filterParams)
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
