package repository

import (
	"gin-starter/internal/model"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type UserRepository interface {
	FindAll() []model.User
	FindById(id uint) (*model.User, error)
	FindByEmail(email string) (*model.User, error)
	Create(user model.User) (*model.User, error)
	Update(user model.User) (*model.User, error)
	Delete(user model.User) error
}

type userRepository struct {
	DB *gorm.DB
}

// Check if UserRepository implements everything in the interface
var _ UserRepository = (*userRepository)(nil)

func NewUserRepository(DB *gorm.DB) UserRepository {
	return &userRepository{DB: DB}
}

func (u *userRepository) FindAll() []model.User {
	var users []model.User

	u.DB.Find(&users)

	return users
}

func (u *userRepository) FindById(id uint) (*model.User, error) {
	var user model.User

	err := u.DB.Where("id = ?", id).First(&user).Error

	return &user, err
}

func (u *userRepository) Create(user model.User) (*model.User, error) {
	err := u.DB.Save(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (u *userRepository) Update(user model.User) (*model.User, error) {
	err := u.DB.Model(&user).Clauses(clause.Returning{}).Updates(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (u *userRepository) Delete(user model.User) error {
	return u.DB.Delete(&user).Error
}

func (u *userRepository) FindByEmail(email string) (*model.User, error) {
	var user model.User

	err := u.DB.Where("email = ?", email).First(&user).Error

	return &user, err
}
