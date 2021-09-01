package aggregate

import (
	"modular-getter/domain/entity/book"
	"modular-getter/domain/entity/user"
)

type Book struct {
	book   book.Book
	author *user.User
}
