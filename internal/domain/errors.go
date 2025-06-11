package domain

import "errors"

var (
	ErrAccountNotFoundByEmail = errors.New("account not found with this id")
	ErrEmailAlreadyExists     = errors.New("this email is already in use")
	ErrPhoneAlreadyExists     = errors.New("this phone is already in use")
	ErrInvalidEmail           = errors.New("invalid email format")
)
