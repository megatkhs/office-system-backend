package usecase

import (
	"backend/model"
	"backend/repository"
)

type EmployeeUsecase interface {
	GetAllEmployees() ([]model.Employee, error)
}

type employeeUsecase struct {
	er repository.EmployeeRepository
}

func NewEmployeeUsecase(er repository.EmployeeRepository) EmployeeUsecase {
	return &employeeUsecase{er}
}

func (eu *employeeUsecase) GetAllEmployees() ([]model.Employee, error) {
	employees := []model.Employee{}

	err := eu.er.GetAllEmployees(&employees)

	return employees, err
}
