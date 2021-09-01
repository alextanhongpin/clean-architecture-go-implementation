// user_builder builds the user representation, allowing the private fields to
// be set from other packages. Particularly useful when constructor starts to
// get huge.

package user

import "github.com/google/uuid"

type Builder struct {
	user Entity
}

func (b Builder) WithID(id uuid.UUID) Builder {
	b.user.id = id
	return b
}

func (b Builder) WithName(name string) Builder {
	b.user.name = name
	return b
}

func (b Builder) Build() Entity {
	return b.user
}
