package user

import "gin-starter/pkg/common"

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
	Email     string `gorm:"column:email" json:"email"`
}

func MapCreateUserDto(dto CreateUserDto) User {
	// TODO: generate a random passoword
	return User{
		FirstName: dto.FirstName,
		LastName:  dto.LastName,
		Email:     dto.Email,
		Password:  "you**must**reset**it**at**the**first**login",
	}
}

type LoginDto struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type RegisterDto struct {
	FirstName            string `json:"firstName"`
	LastName             string `json:"lastName"`
	Email                string `json:"email" binding:"required"`
	Password             string `json:"password" binding:"required"`
	PasswordConfirmation string `json:"passwordConfirmation" binding:"required, eq|eqfield=Password"`
}

func MapRegisterDto(dto RegisterDto) User {
	// TODO: generate a random passoword
	return User{
		FirstName: dto.FirstName,
		LastName:  dto.LastName,
		Email:     dto.Email,
		Password:  dto.Password,
	}
}
