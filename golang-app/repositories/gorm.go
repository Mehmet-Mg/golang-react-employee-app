package repositories

import (
	"context"
	"example/employee-app/gormstore"
	"example/employee-app/models"

	"gorm.io/gorm"
)

type GormEmployeeRepository struct {
	db *gorm.DB
}

func NewGormEmployeeRepository(db *gorm.DB) *GormEmployeeRepository {
	return &GormEmployeeRepository{
		db: db,
	}
}

func (r *GormEmployeeRepository) CreateEmployee(ctx context.Context, employee *models.Employee) (*models.Employee, error) {
	return gormstore.CreateEmployee(ctx, r.db, employee)
}

func (r *GormEmployeeRepository) GetAllEmployees(ctx context.Context) ([]models.Employee, error) {
	return gormstore.GetAllEmployees(ctx, r.db)
}

func (r *GormEmployeeRepository) GetEmployeeById(ctx context.Context, id uint64) (*models.Employee, error) {
	return gormstore.GetEmployeeById(ctx, r.db, id)
}

func (r *GormEmployeeRepository) UpdateEmployee(ctx context.Context, id uint64, employee *models.Employee) (*models.Employee, error) {
	return gormstore.UpdateEmployee(ctx, r.db, id, employee)
}

func (r *GormEmployeeRepository) DeleteEmployee(ctx context.Context, id uint64) error {
	return gormstore.Delete(ctx, r.db, id)
}
