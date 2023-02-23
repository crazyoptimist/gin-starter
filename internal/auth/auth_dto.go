package auth

type LoginDto struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type RegisterDto struct {
	FirstName            string `json:"firstName"`
	LastName             string `json:"lastName"`
	Address              string `json:"address"`
	Email                string `json:"email" binding:"required"`
	Password             string `json:"password" binding:"required"`
	PasswordConfirmation string `json:"passwordConfirmation" binding:"required, eq|eqfield=Password"`
}
