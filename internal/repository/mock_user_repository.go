package repository

import (
	"errors"

	"gin-starter/internal/model"
	"gin-starter/pkg/common"
	"gin-starter/pkg/utils"
)

type mockUserRepository struct {
	records []model.User
}

var _ UserRepository = (*mockUserRepository)(nil)

func NewMockUserRepository() UserRepository {
	return &mockUserRepository{
		records: []model.User{
			{BaseModel: common.BaseModel{ID: 1}, FirstName: "John", LastName: "Smith", Email: "john.smith@gmail.com", Password: "password"},
			{BaseModel: common.BaseModel{ID: 2}, FirstName: "Ben", LastName: "Doe", Email: "ben.doe@gmail.com", Password: "password"},
		},
	}
}

func (r *mockUserRepository) FindAll() []model.User {
	return r.records
}

func (r *mockUserRepository) FindById(id uint) (*model.User, error) {
	for _, record := range r.records {
		if record.ID == id {
			return &record, nil
		}
	}
	return nil, errors.New("User not found with given ID")
}

func (r *mockUserRepository) FindByEmail(email string) (*model.User, error) {
	for _, record := range r.records {
		if record.Email == email {
			return &record, nil
		}
	}
	return nil, errors.New("User not found with given Email")
}

func (r *mockUserRepository) Create(user model.User) (*model.User, error) {
	r.records = append(r.records, user)
	return &user, nil
}

func (r *mockUserRepository) Update(user model.User) (*model.User, error) {
	for i, record := range r.records {
		if record.ID == user.ID {
			r.records[i] = user
			return &user, nil
		}
	}
	return &user, nil
}

func (r *mockUserRepository) Delete(user model.User) error {
	utils.RemoveAt(r.records, int(user.ID))
	return nil
}
