package customerrors

import "errors"

var (
	ErrInvalidPhoneNumber = errors.New("invalid phone number")
	ErrInvalidEmail       = errors.New("invalid email")
	ErrInvalidPassword    = errors.New("invalid password")
	ErrInvalidName        = errors.New("invalid name")
)
