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

func (u *userRepository) FindAll(
	paginationParam utils.PaginationParam,
	sortParams []utils.SortParam,
	filterParams []utils.FilterParam,
) []model.User {
	var users []model.User

	var orderQuery string

	for _, sortParam := range sortParams {
		if sortParam.FieldName == "email" {
			orderQuery += " email " + sortParam.Order
		}
		if sortParam.FieldName == "firstName" {
			orderQuery += " ethnicity " + sortParam.Order
		}
		if sortParam.FieldName == "lastName" {
			orderQuery += " last_name " + sortParam.Order
		}
	}

	u.DB.Order(
		orderQuery,
	).Limit(
		paginationParam.Limit,
	).Offset(
		paginationParam.Offset,
	).Find(&users)

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
