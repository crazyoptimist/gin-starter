package user

import "gin-starter/cmd/api/config"

type UserDAO struct{}

func NewUserDAO() *UserDAO {
	return &UserDAO{}
}

func (dao *UserDAO) Get(id uint) (*User, error) {
	var user User

	err := config.Config.DB.Where("id = ?", id).
		First(&user).
		Error

	return &user, err
}
