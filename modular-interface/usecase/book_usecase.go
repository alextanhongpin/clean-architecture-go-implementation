package usecase

import (
	"errors"
	"modular-interface/domain/aggregate/books"
)

type BookUsecase struct {
	repo books.Repository
}

func (uc BookUsecase) Create() error {
	return errors.New("not implemented")
}
