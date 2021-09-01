package model

import "github.com/google/uuid"

type Book struct {
	ID    int64
	Title string

	// There's no issue with circular reference, since all models belongs to the
	// same package.
	// However, it might be better to just reference the author by user id rather
	// then embedding the User struct.
	// The embedding could be done in the aggregate layer instead.
	Author   *User
	AuthorID uuid.UUID
}
