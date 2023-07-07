package models

import (
	"context"
)

type EmployeeRepository interface {
	Create(ctx context.Context, e *Employee) (*Employee, error)
	All(ctx context.Context) ([]Employee, error)
	GetById(ctx context.Context, id int) (*Employee, error)
	Update(ctx context.Context, id int, e *Employee) (*Employee, error)
	Delete(ctx context.Context, id int) error
}
