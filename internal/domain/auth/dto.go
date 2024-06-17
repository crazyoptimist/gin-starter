package auth

type LoginDto struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginResponse struct {
	AccessToken  string  `json:"accessToken"`
	RefreshToken string  `json:"refreshToken"`
	ExpiresIn    float64 `json:"expiresIn"`
}

type RefreshDto struct {
	RefreshToken string `json:"refreshToken"`
}

type RegisterDto struct {
	FirstName            string `json:"firstName"`
	LastName             string `json:"lastName"`
	Email                string `json:"email" binding:"required"`
	Password             string `json:"password" binding:"required"`
	PasswordConfirmation string `json:"passwordConfirmation" binding:"required"`
}
