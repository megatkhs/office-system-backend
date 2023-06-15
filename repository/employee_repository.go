package repository

import (
	"backend/model"

	"gorm.io/gorm"
)

type EmployeeRepository interface {
	GetAllEmployees(employees *[]model.Employee) error
	GetEmployeeById(employee *model.Employee, id uint) error
}

type employeeRepository struct {
	db *gorm.DB
}

func NewEmployeeRepository(db *gorm.DB) EmployeeRepository {
	return &employeeRepository{db}
}

func (er *employeeRepository) GetAllEmployees(employees *[]model.Employee) error {
	err := er.db.Model(&model.Employee{}).Find(employees).Error
	return err
}

func (er *employeeRepository) GetEmployeeById(employee *model.Employee, id uint) error {
	err := er.db.Preload("Contacts").First(employee, id).Error
	return err
}
