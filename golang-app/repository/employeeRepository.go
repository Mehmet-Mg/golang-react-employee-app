package repository

import (
	"example/employee-app/models"

	"gorm.io/gorm"
)

type EmployeRepository struct {
	db *gorm.DB
}

func NewEmployeeRepository(db *gorm.DB) *EmployeRepository {
	return &EmployeRepository{
		db: db,
	}
}

func (r *EmployeRepository) Migrate() error {
	return r.db.AutoMigrate(&models.Employee{})
}

func (r *EmployeRepository) Create(employee models.Employee) (*models.Employee, error) {
	err := r.db.Create(employee).Error

	if err != nil {
		return nil, ErrDuplicate
	}

	return &employee, nil
}

func (r *EmployeRepository) All() ([]models.Employee, error) {
	var employees []models.Employee

	if err := r.db.Find(&employees).Error; err != nil {
		return nil, ErrNotExist
	}

	return employees, nil
}

func (r *EmployeRepository) GetById(id uint64) (*models.Employee, error) {
	var employee models.Employee
	if err := r.db.First(&employee, id).Error; err != nil {
		return nil, err
	}

	return &employee, nil
}

func (r *EmployeRepository) Update(id uint64, updated models.Employee) (*models.Employee, error) {
	if err := r.db.Model(&updated).Updates(updated).Error; err != nil {
		return nil, err
	}
	return &updated, nil
}

func (r *EmployeRepository) Delete(id uint64) error {
	var employee models.Employee
	if err := r.db.Delete(&employee, id).Error; err != nil {
		return err
	}

	return nil
}
