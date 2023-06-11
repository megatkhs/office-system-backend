package repository

import (
	"backend/model"

	"gorm.io/gorm"
)

type EmployeeRepository interface {
	GetAllEmployees(employees *[]model.Employee) error
}

type employeeRepository struct {
	db *gorm.DB
}

func NewEmployeeRepository(db *gorm.DB) EmployeeRepository {
	return &employeeRepository{db}
}

func (er *employeeRepository) GetAllEmployees(employees *[]model.Employee) error {
	err := er.db.Model(&model.Employee{}).Preload("Contacts").Find(employees).Error
	return err
}
