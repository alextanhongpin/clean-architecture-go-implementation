package user

import (
	"context"

	"github.com/google/uuid"
)

type Repository interface {
	Find(ctx context.Context, id uuid.UUID) (*User, error)
	Create(ctx context.Context, user CreateUserDto) (*User, error)
	Update(ctx context.Context, user UpdateUserDto) (*User, error)
	Delete(ctx context.Context, id uuid.UUID) (*User, error)
}

// Interface vs Struct for DTO.
// Interface allows us to pass the full struct, while selectively selecting the
// fields to be updated.
// Struct requires us to map the representation of the user to another data structure.

type CreateUserDto interface {
	Name() string
}

type UpdateUserDto struct {
	ID   uuid.UUID
	Name string
}
