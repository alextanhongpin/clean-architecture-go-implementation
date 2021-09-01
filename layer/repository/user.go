package repository

import (
	"context"

	"layer/model"

	"github.com/google/uuid"
)

type UserRepository interface {
	Find(ctx context.Context, id uuid.UUID) (*model.User, error)
	Create(ctx context.Context, user model.User) (*model.User, error)
	Update(ctx context.Context, user model.User) (*model.User, error)
	Delete(ctx context.Context, id uuid.UUID) (*model.User, error)
}
