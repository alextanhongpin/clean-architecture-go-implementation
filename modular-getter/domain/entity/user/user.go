package user

import (
	"errors"

	"github.com/google/uuid"
)

type User struct {
	id   uuid.UUID
	name string
}

func (u User) ID() uuid.UUID { return u.id }
func (u User) Name() string  { return u.name }

// SetName sets the name.
func (u *User) SetName(name string) error {
	if name == "" {
		return errors.New(`user: name is required`)
	}

	u.name = name
	return nil
}

// WithName sets the name and return a copy of the user with the new name.
// Wither is preferred to setter.
func (u User) WithName(name string) (User, error) {
	if name == "" {
		return u, errors.New(`user: name is required`)
	}

	u.name = name
	return u, nil
}
