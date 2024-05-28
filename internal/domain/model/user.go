package model

type User struct {
	Common
	FirstName string `gorm:"column:first_name" json:"firstName"`
	LastName  string `gorm:"column:last_name" json:"lastName"`
	Email     string `gorm:"column:email;unique" json:"email"`
	Password  string `gorm:"column:password" json:"-"`
}
