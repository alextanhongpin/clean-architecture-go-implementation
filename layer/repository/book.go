package repository

import (
	"context"

	"layer/model"

	"github.com/google/uuid"
)

type BookRepository interface {
	Find(ctx context.Context, id uuid.UUID) (*model.Book, error)
	Create(ctx context.Context, book model.Book) (*model.Book, error)
	Update(ctx context.Context, book model.Book) (*model.Book, error)
	Delete(ctx context.Context, id uuid.UUID) (*model.Book, error)
}
