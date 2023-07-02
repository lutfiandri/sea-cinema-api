package contract

import "time"

type RegisterRequest struct {
	Username  string    `json:"username" validate:"required,min=3"`
	Password  string    `json:"password" validate:"required,min=8"`
	Name      string    `json:"name" validate:"required"`
	BirthDate time.Time `json:"birth_date" validate:"required"`
}

type LoginRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type RegisterResponse struct {
	AccessToken string `json:"access_token"`
}

type LoginResponse struct {
	AccessToken string `json:"access_token"`
}
