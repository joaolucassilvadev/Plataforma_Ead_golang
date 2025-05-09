package entities

import (
	"errors"

	"github.com/google/uuid"
)

type User struct {
	BaseEntity
	Name     string
	Email    string
	Phone    string
	Password string
}

func NewUser(name, email, phone, password string) (*User, error) {
	user := &User{
		Name:     name,
		Email:    email,
		Phone:    phone,
		Password: password,
	}
	user.Id = uuid.New()

	if err := user.Validate(); err != nil {
		return nil, err
	}
	return user, nil
}

func (u *User) Validate() error {
	if u.Name == "" {
		return errors.New("user name invalid")
	}
	if u.Email == "" {
		return errors.New("user email invalid")
	}
	if u.Phone == "" {
		return errors.New("user phone invalid")
	}
	if u.Password == "" || len(u.Password) < 8 {
		return errors.New("Password must be at least 8 characters long")
	}
	return nil
}
