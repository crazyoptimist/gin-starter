package repository

import (
	"gin-starter/internal/domain/model"
	"gin-starter/internal/domain/user"
	"gin-starter/pkg/common"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type userRepository struct {
	DB *gorm.DB
}

var _ user.UserRepository = (*userRepository)(nil)

func NewUserRepository(DB *gorm.DB) *userRepository {
	return &userRepository{DB: DB}
}

func (u *userRepository) FindAll(
	paginationParam common.PaginationParam,
	sortParams []common.SortParam,
	filterParams []common.FilterParam,
) ([]model.User, int64, error) {
	query := u.DB.Model(&model.User{})

	// Filter
	filterQuery := make(map[string]interface{})

	for _, filterParam := range filterParams {
		if filterParam.FieldName == "email" {
			filterQuery["email"] = filterParam.Value
		}
		if filterParam.FieldName == "firstName" {
			filterQuery["first_name"] = filterParam.Value
		}
		if filterParam.FieldName == "lastName" {
			filterQuery["last_name"] = filterParam.Value
		}
	}

	query = query.Where(filterQuery)

	// Count
	var totalCount int64

	if err := query.Count(&totalCount).Error; err != nil {
		return nil, 0, err
	}

	// Sort
	for _, sortParam := range sortParams {
		if sortParam.FieldName == "email" {
			query = query.Order("email " + sortParam.Order)
		}
		if sortParam.FieldName == "firstName" {
			query = query.Order("first_name " + sortParam.Order)
		}
		if sortParam.FieldName == "lastName" {
			query = query.Order("last_name " + sortParam.Order)
		}
	}

	// Pagination
	query = query.Limit(
		paginationParam.Limit,
	).Offset(
		paginationParam.Offset,
	)

	var users []model.User

	if err := query.Find(&users).Error; err != nil {
		return nil, 0, err
	}

	return users, totalCount, nil
}

func (u *userRepository) FindById(id int) (*model.User, error) {
	var user model.User

	err := u.DB.Where("id = ?", id).First(&user).Error

	return &user, err
}

func (u *userRepository) FindByEmail(email string) (*model.User, error) {
	var user model.User

	err := u.DB.Where("email = ?", email).First(&user).Error

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

func (u *userRepository) UpdateRefreshToken(userId int, refreshToken string) error {
	err := u.DB.Model(
		&model.User{},
	).Where(
		"id = ?",
		userId,
	).Updates(
		map[string]interface{}{
			"RefreshToken": refreshToken,
		},
	).Error
	if err != nil {
		return err
	}

	return nil
}

func (u *userRepository) Delete(user model.User) error {
	return u.DB.Delete(&user).Error
}
