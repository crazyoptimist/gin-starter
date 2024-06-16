package user

import (
	"errors"

	"gin-starter/internal/domain/model"
	"gin-starter/pkg/common"
	"gin-starter/pkg/utils"
)

type userRepositoryStub struct {
	records []model.User
}

var _ UserRepository = (*userRepositoryStub)(nil)

func NewUserRepositoryStub() *userRepositoryStub {
	return &userRepositoryStub{
		records: []model.User{
			{
				Common:    model.Common{ID: 1},
				FirstName: "John",
				LastName:  "Smith",
				Email:     "john.smith@gmail.com",
				Password:  "password",
			},
			{
				Common:    model.Common{ID: 2},
				FirstName: "Ben",
				LastName:  "Doe",
				Email:     "ben.doe@gmail.com",
				Password:  "password",
			},
		},
	}
}

func (r *userRepositoryStub) FindAll(
	paginationParam common.PaginationParam,
	sortParams []common.SortParam,
	filterParams []common.FilterParam,
) ([]model.User, int64, error) {
	return r.records, 2, nil
}

func (r *userRepositoryStub) FindById(id int) (*model.User, error) {
	for _, record := range r.records {
		if record.ID == id {
			return &record, nil
		}
	}
	return nil, errors.New("User not found with the given ID")
}

func (r *userRepositoryStub) FindByEmail(email string) (*model.User, error) {
	for _, record := range r.records {
		if record.Email == email {
			return &record, nil
		}
	}
	return nil, errors.New("User not found with the given Email")
}

func (r *userRepositoryStub) Create(user model.User) (*model.User, error) {
	r.records = append(r.records, user)
	return &user, nil
}

func (r *userRepositoryStub) Update(user model.User) (*model.User, error) {
	for i, record := range r.records {
		if record.ID == user.ID {
			r.records[i] = user
			return &user, nil
		}
	}
	return &user, nil
}

func (r *userRepositoryStub) Delete(user model.User) error {
	utils.RemoveByIndex(r.records, int(user.ID))
	return nil
}
