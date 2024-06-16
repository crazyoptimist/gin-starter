package user

import (
	"github.com/google/uuid"

	"gin-starter/internal/domain/model"
	"gin-starter/pkg/utils"
)

type CreateUserDto struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email" binding:"required,email"`
}

func MapCreateUserDto(dto *CreateUserDto) model.User {
	password := uuid.NewString()
	hashedPassword, _ := utils.HashPassword(password)

	return model.User{
		FirstName: dto.FirstName,
		LastName:  dto.LastName,
		Email:     dto.Email,
		Password:  hashedPassword,
	}
}

type UpdateUserDto struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email" binding:"omitempty,email"`
	Password  string `json:"password"`
}

func MapUpdateUserDto(dto *UpdateUserDto, id int) model.User {
	var hashedPassword string
	if dto.Password != "" {
		hashedPassword, _ = utils.HashPassword(dto.Password)
	}

	return model.User{
		Common:    model.Common{ID: id},
		FirstName: dto.FirstName,
		LastName:  dto.LastName,
		Email:     dto.Email,
		Password:  hashedPassword,
	}
}
