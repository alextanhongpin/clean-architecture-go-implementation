package user

import (
	"context"

	"github.com/google/uuid"
)

type Repository interface {
	Find(ctx context.Context, id uuid.UUID) (*User, error)
	Create(ctx context.Context, user User) (*User, error)
	Update(ctx context.Context, user User) (*User, error)
	Delete(ctx context.Context, id uuid.UUID) (*User, error)
}
