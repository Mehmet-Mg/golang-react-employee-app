package gormstore

import (
	"context"
	"example/employee-app/models"

	"gorm.io/gorm"
)

const (
	EmployeeCollection = "employees"
)

func CreateEmployee(ctx context.Context, db *gorm.DB, employee *models.Employee) (*models.Employee, error) {
	err := db.WithContext(ctx).Create(employee).Error

	if err != nil {
		return nil, err
	}

	return employee, nil
}

func GetAllEmployees(ctx context.Context, db *gorm.DB) ([]models.Employee, error) {
	var entities []models.Employee

	if err := db.WithContext(ctx).Find(&entities).Error; err != nil {
		return nil, err
	}

	return entities, nil
}

func GetEmployeeById(ctx context.Context, db *gorm.DB, id uint64) (*models.Employee, error) {
	var entity models.Employee
	if err := db.WithContext(ctx).First(&entity, id).Error; err != nil {
		return nil, err
	}

	return &entity, nil
}

func UpdateEmployee(ctx context.Context, db *gorm.DB, id uint64, entity *models.Employee) (*models.Employee, error) {
	if err := db.Model(&entity).WithContext(ctx).Updates(entity).Error; err != nil {
		return nil, err
	}
	return entity, nil
}

func Delete(ctx context.Context, db *gorm.DB, id uint64) error {
	var entity models.Employee
	if err := db.WithContext(ctx).Delete(entity, id).Error; err != nil {
		return err
	}

	return nil
}
