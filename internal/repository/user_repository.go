package repository

import (
	"clean-go/internal/entity"
	"errors"
)

var ErrNotFound = errors.New("not found")

type UserRepository interface {
	All() ([]*entity.User, error)
	GetByEmail(string) (*entity.User, error)
	GetById(string) (*entity.User, error)
	Create(*entity.User) error
}
