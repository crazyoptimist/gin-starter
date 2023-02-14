package user

import "gin-starter/cmd/api/config"

// UserDAO persists user data in database
type UserDAO struct{}

// NewUserDAO creates a new UserDAO
func NewUserDAO() *UserDAO {
	return &UserDAO{}
}

// Get does the actual query to database, if user with specified id is not found error is returned
func (dao *UserDAO) Get(id uint) (*User, error) {
	var user User

	// Query Database here...

	//user = User{
	//	Model: Model{ID: 1},
	//	FirstName: "John",
	//	LastName: "Doe",
	//	Address: "Some Address",
	//	Email: "ninja@example.com"}

	// if using Gorm:
	err := config.Config.DB.Where("id = ?", id).
		First(&user).
		Error

	return &user, err
}
