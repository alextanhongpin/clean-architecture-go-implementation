// books aggregate can have it's own repository.
// The repository usually returns the composite entities.

package books

type Repository interface {
	Find() []Book
}
