package user

import (
	"gin-starter/internal/domain/model"
	"gin-starter/pkg/common"
)

type UserRepository interface {
	FindAll(
		paginationParam common.PaginationParam,
		sortParams []common.SortParam,
		filterParams []common.FilterParam,
	) ([]model.User, int64, error)
	FindById(id uint) (*model.User, error)
	FindByEmail(email string) (*model.User, error)
	Create(user model.User) (*model.User, error)
	Update(user model.User) (*model.User, error)
	Delete(user model.User) error
}

type UserService struct {
	UserRepository UserRepository
}

func NewUserService(userRepository UserRepository) *UserService {
	return &UserService{UserRepository: userRepository}
}

// FindAll returns the records found, total count, and error
func (u *UserService) FindAll(
	paginationParam common.PaginationParam,
	sortParams []common.SortParam,
	filterParams []common.FilterParam,
) ([]model.User, int64, error) {
	return u.UserRepository.FindAll(paginationParam, sortParams, filterParams)
}

func (u *UserService) FindById(id uint) (*model.User, error) {
	return u.UserRepository.FindById(id)
}

func (u *UserService) Create(createUserDto *CreateUserDto) (*model.User, error) {
	return u.UserRepository.Create(MapCreateUserDto(createUserDto))
}

func (u *UserService) Update(updateUserDto *UpdateUserDto, id uint) (*model.User, error) {
	return u.UserRepository.Update(MapUpdateUserDto(updateUserDto, id))
}

func (u *UserService) Delete(id uint) error {
	return u.UserRepository.Delete(model.User{Common: model.Common{ID: id}})
}
