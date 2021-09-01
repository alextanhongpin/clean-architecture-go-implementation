package user

import (
	"errors"
	"strings"
)

// Service is a struct ...
type Service struct {
}

// Service should be better off as functions.
func NewUserName(name string) (string, error) {
	name = strings.TrimSpace(name)
	if name == "" {
		return name, errors.New("user: name is required")
	}
	return name, nil
}
