package models

import "errors"

var ErrUserNotFoundOrWrongPassword = errors.New("user not found or invalid password")

type (
	Me struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	}

	LoginRequest struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	CreateUserRequest struct {
		Email           string `json:"email"`
		Name            string `json:"name"`
		Password        string `json:"password"`
		ConfirmPassword string `json:"confirmPassword"`
	}
)

func (c CreateUserRequest) Validate() error {
	if c.Email == "" {
		return errors.New("email field must be provided")
	}
	if c.Name == "" {
		return errors.New("name field must be provided")
	}
	if c.Password == "" {
		return errors.New("password field must be provided")
	}
	if c.ConfirmPassword == "" {
		return errors.New("confirmPassword field must be provided")
	}

	if c.Password != c.ConfirmPassword {
		return errors.New("password and confirmPassword fields doens't match")
	}
	return nil
}
