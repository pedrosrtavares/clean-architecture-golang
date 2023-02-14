package entity

import (
	"clean-architecture-golang/internal/customerrors"
	"clean-architecture-golang/internal/validator"
)

type User struct {
	ID          ID
	FirstName   string
	LastName    string
	Email       string
	Password    string
	PhoneNumber string
}

func NewUser(firstName, lastName, email, password, phoneNumber string) (*User, error) {
	u := &User{
		ID:          NewID(),
		FirstName:   firstName,
		LastName:    lastName,
		Email:       email,
		Password:    password,
		PhoneNumber: phoneNumber,
	}

	err := u.Validate()
	if err != nil {
		return nil, err
	}

	return u, nil
}

func (u *User) Validate() error {
	if u.FirstName == "" || u.LastName == "" {
		return customerrors.ErrInvalidName
	}

	if u.Email == "" {
		return validator.IsValidEmail(u.Email)
	}

	if u.Password == "" {
		return validator.IsValidPassword(u.Password)
	}

	if u.PhoneNumber == "" {
		return validator.IsValidPhone(u.PhoneNumber)
	}

	return nil
}
