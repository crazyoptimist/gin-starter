package user

import (
	"gin-starter/pkg/common"
	"gin-starter/pkg/utils"
)

type User struct {
	common.Model
	FirstName string `gorm:"column:first_name" json:"firstName"`
	LastName  string `gorm:"column:last_name" json:"lastName"`
	Email     string `gorm:"column:email" json:"email"`
	Password  string `gorm:"column:password" json:"-"`
}

type CreateUserDto struct {
	FirstName string `gorm:"column:first_name" json:"firstName"`
	LastName  string `gorm:"column:last_name" json:"lastName"`
	Email     string `gorm:"column:email" json:"email" binding:"required,email"`
}

func MapCreateUserDto(dto *CreateUserDto) User {
	// TODO: generate a random passoword
	password, _ := utils.HashPassword("must**reset")

	return User{
		FirstName: dto.FirstName,
		LastName:  dto.LastName,
		Email:     dto.Email,
		Password:  password,
	}
}

type UpdateUserDto struct {
	FirstName string `gorm:"column:first_name" json:"firstName"`
	LastName  string `gorm:"column:last_name" json:"lastName"`
	Email     string `gorm:"column:email" json:"email"`
}

func MapUpdateUserDto(dto *UpdateUserDto, id uint) User {
	return User{
		Model:     common.Model{ID: id},
		FirstName: dto.FirstName,
		LastName:  dto.LastName,
		Email:     dto.Email,
	}
}
