package user

import "gin-starter/pkg/common"

type UserService struct {
	UserRepository IUserRepository
}

func NewUserService(repository IUserRepository) UserService {
	return UserService{UserRepository: repository}
}

func (u *UserService) FindAll() []User {
	return u.UserRepository.FindAll()
}

func (u *UserService) FindById(id uint) (*User, error) {
	return u.UserRepository.FindById(id)
}

func (u *UserService) Create(dto *CreateUserDto) (*User, error) {
	return u.UserRepository.Create(MapCreateUserDto(dto))
}

func (u *UserService) Update(dto *UpdateUserDto, id uint) (*User, error) {
	return u.UserRepository.Update(MapUpdateUserDto(dto, id))
}

func (u *UserService) Delete(id uint) error {
	return u.UserRepository.Delete(User{Model: common.Model{ID: id}})
}
