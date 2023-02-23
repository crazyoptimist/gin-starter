package user

import (
	"errors"
	"gin-starter/pkg/common"
	"gin-starter/pkg/utils"
)

type mockUserRepository struct {
	records []User
}

var _ IUserRepository = (*mockUserRepository)(nil)

func newMockUserRepository() *mockUserRepository {
	return &mockUserRepository{
		records: []User{
			{Model: common.Model{ID: 1}, FirstName: "John", LastName: "Smith", Email: "john.smith@gmail.com", Password: "password"},
			{Model: common.Model{ID: 2}, FirstName: "Ben", LastName: "Doe", Email: "ben.doe@gmail.com", Password: "password"},
		},
	}
}

func (r *mockUserRepository) Save(user User) (*User, error) {
	if user.ID != 0 {
		for i, record := range r.records {
			if record.ID == user.ID {
				r.records[i] = user
				return &user, nil
			}
		}
	}
	r.records = append(r.records, user)
	return &user, nil
}

func (r *mockUserRepository) FindAll() []User {
	return r.records
}

func (r *mockUserRepository) FindById(id uint) (*User, error) {
	for _, record := range r.records {
		if record.ID == id {
			return &record, nil
		}
	}
	return nil, errors.New("not found")
}

func (r *mockUserRepository) Delete(user User) error {
	utils.RemoveAt(r.records, int(user.ID))
	return nil
}
