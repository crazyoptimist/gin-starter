package dto

import (
	"gin-starter/internal/model"
	"gin-starter/pkg/common"
	"gin-starter/pkg/utils"
)

type CreateUserDto struct {
	FirstName string `gorm:"column:first_name" json:"firstName"`
	LastName  string `gorm:"column:last_name" json:"lastName"`
	Email     string `gorm:"column:email" json:"email" binding:"required,email"`
}

func MapCreateUserDto(dto *CreateUserDto) model.User {
	// TODO: generate a random passoword
	password, _ := utils.HashPassword("must**reset")

	return model.User{
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

func MapUpdateUserDto(dto *UpdateUserDto, id uint) model.User {
	return model.User{
		BaseModel: common.BaseModel{ID: id},
		FirstName: dto.FirstName,
		LastName:  dto.LastName,
		Email:     dto.Email,
	}
}
