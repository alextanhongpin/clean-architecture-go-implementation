package aggregate

import (
	"modular-interface/domain/entity/book"
	"modular-interface/domain/entity/user"
)

type Book struct {
	book   book.Book
	author *user.User
}
