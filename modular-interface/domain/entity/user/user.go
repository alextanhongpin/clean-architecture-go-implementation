package user

import (
	"errors"

	"github.com/google/uuid"
)

// Query defines the read-only part of the entity.
type Query interface {
	ID() uuid.UUID
	Name() string
}

// Mutator defines the write-only part of the entity.
type Mutator interface {
	WithName(name string) (User, error)
}

// User ...
// Why do we need a separate Query and Mutator interface?
// At the API layer, we want to return only types that are understood by
// the client, but do not want to allow modification to the entity anymore.
// Hence, returning a read-only contract would make more sense then exposing
// the whole entity.
type User interface {
	Query
	Mutator
}

var _ User = (*Entity)(nil)

type Entity struct {
	id   uuid.UUID
	name string
}

func (e Entity) ID() uuid.UUID { return e.id }
func (e Entity) Name() string  { return e.name }

func (e Entity) WithName(name string) (User, error) {
	if name == "" {
		return e, errors.New("user: name is required")
	}

	e.name = name
	return e, nil
}
