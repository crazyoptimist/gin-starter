package user

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type IUserRepository interface {
	FindAll() []User
	FindById(id uint) (*User, error)
	Create(user User) (*User, error)
	Update(user User) (*User, error)
	Delete(user User) error
}

// Check if the implementation satisfies all methods of interface
var _ IUserRepository = (*UserRepository)(nil)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(DB *gorm.DB) *UserRepository {
	return &UserRepository{DB: DB}
}

func (u *UserRepository) FindAll() []User {
	var users []User

	u.DB.Find(&users)

	return users
}

func (u *UserRepository) FindById(id uint) (*User, error) {
	var user User

	err := u.DB.Where("id = ?", id).First(&user).Error

	return &user, err
}

func (u *UserRepository) Create(user User) (*User, error) {
	err := u.DB.Save(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (u *UserRepository) Update(user User) (*User, error) {
	err := u.DB.Model(&user).Clauses(clause.Returning{}).Updates(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (u *UserRepository) Delete(user User) error {
	return u.DB.Delete(&user).Error
}

func (u *UserRepository) FindByEmail(email string) (*User, error) {
	var user User

	err := u.DB.Where("email = ?", email).Find(&user).Error

	return &user, err
}
