package repository

import (
	"gin-starter/internal/model"
	"gin-starter/pkg/utils"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type userRepository struct {
	DB *gorm.DB
}

func NewUserRepository(DB *gorm.DB) *userRepository {
	return &userRepository{DB: DB}
}

func (u *userRepository) FindAll(queryParams utils.QueryParams) []model.User {
	var users []model.User

	u.DB.Limit(queryParams.Limit).Offset(queryParams.Offset).Find(&users)

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
