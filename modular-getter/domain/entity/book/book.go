package book

import "github.com/google/uuid"

type Book struct {
	id       int64
	title    string
	authorID uuid.UUID
}

func (b Book) ID() int64           { return b.id }
func (b Book) Title() string       { return b.title }
func (b Book) AuthorID() uuid.UUID { return b.authorID }
