package user

import (
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
	validName, err := NewUserName(name)
	if err != nil {
		return e, err
	}
	e.name = validName
	return e, nil
}

// SimpleQuery returns an even smaller scope of the field.
// Using interface is powerful, it allows you to restrict fields without
// modifying structs.
// This can even be returned at the application layer (usecase) if we want an
// even more simpler representation instead of the full domain entity.
type SimpleQuery interface {
	Name() string
}
