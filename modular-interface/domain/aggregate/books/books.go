// books aggregate represents a collection of books with authors.

package books

import (
	"modular-interface/domain/entity/book"
	"modular-interface/domain/entity/user"
)

type Book struct {
	book   book.Book
	author user.User
}

// aggregate may operate at collection level, or may even return the summarized
// result from the repository.
type BookList struct {
	books []Book
}

func (b BookList) Unique() bool {
	panic("not implemented")
	return false
}
