package params

import "go_sample_login_register/validators"

type LoginRequest struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type RegisterRequest struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func (request LoginRequest) Validate() error {
	return validators.ValidateInputs(request)
}

func (request RegisterRequest) Validate() error {
	return validators.ValidateInputs(request)
}

type LoginResponse struct {
	AuthToken string `json:"auth_token"`
}

type RegisterResponse struct {
	AuthToken string `json:"auth_token"`
	Email     string `json:"email"`
}

type MeResponse struct {
	Email string `json:"email" validate:"required"`
}
