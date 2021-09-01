package book

import "github.com/google/uuid"

type Book struct {
	ID       int64
	Title    string
	AuthorID uuid.UUID
}
