package repository

import (
	"errors"
)

var (
	ErrDuplicate    = errors.New("record already exists")
	ErrNotExist     = errors.New("row does not exist")
	ErrUpdateFailed = errors.New("update failed")
	ErrDeleteFailed = errors.New("delete failed")
)

type Repository[T any] interface {
	Migrate() error
	Create(entity T) (*T, error)
	All() ([]T, error)
	GetById(id uint) (*T, error)
	Update(id uint, updated T) (*T, error)
	Delete(id uint) error
}
