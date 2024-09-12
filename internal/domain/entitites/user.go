package entities

import (
	"errors"
)

type User struct {
	BaseEntity
	Username string
}

func NewUser(username string) (*User, error) {
	user := &User{
		BaseEntity: *NewBaseEntity(),
		Username:   username,
	}
	if err := user.Validate(); err != nil {
		return nil, err
	}
	return user, nil
}

func (u *User) Validate() error {
	if u.Username == "" {
		return errors.New("username cannot be empty")
	}
	return nil
}
