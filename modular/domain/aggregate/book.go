package aggregate

import (
	"modular/domain/entity/book"
	"modular/domain/entity/user"
)

type Book struct {
	book   book.Book
	author *user.User
}
