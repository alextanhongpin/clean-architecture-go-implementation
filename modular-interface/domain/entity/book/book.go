package book

import "github.com/google/uuid"

type Entity struct {
	ID       int64
	Title    string
	AuthorID uuid.UUID
}

type Book interface {
	ID() int64
	Title() string
	AuthorID() uuid.UUID
}
