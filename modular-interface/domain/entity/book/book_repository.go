package book

import (
	"context"

	"github.com/google/uuid"
)

type Repository interface {
	Find(ctx context.Context, id uuid.UUID) (*Book, error)
	Create(ctx context.Context, book Book) (*Book, error)
	Update(ctx context.Context, book Book) (*Book, error)
	Delete(ctx context.Context, id uuid.UUID) (*Book, error)
}
