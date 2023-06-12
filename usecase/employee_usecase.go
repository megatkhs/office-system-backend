package usecase

import (
	"backend/model"
	"backend/repository"
)

type EmployeeUsecase interface {
	GetAllEmployees() ([]model.EmployeeResponse, error)
}

type employeeUsecase struct {
	er repository.EmployeeRepository
}

func NewEmployeeUsecase(er repository.EmployeeRepository) EmployeeUsecase {
	return &employeeUsecase{er}
}

func (eu *employeeUsecase) GetAllEmployees() ([]model.EmployeeResponse, error) {
	employees := []model.Employee{}

	if err := eu.er.GetAllEmployees(&employees); err != nil {
		return nil, err
	}

	employeesRes := make([]model.EmployeeResponse, len(employees))
	for i, v := range employees {
		employeesRes[i] = model.EmployeeResponse{
			ID:           v.ID,
			FirstName:    v.FirstName,
			LastName:     v.LastName,
			Birthday:     v.Birthday,
			DepartmentID: v.DepartmentID,
			CreatedAt:    v.CreatedAt,
			UpdatedAt:    v.UpdatedAt,
			Contacts:     make([]model.ContactResponse, len(v.Contacts)),
		}

		for ci, cv := range v.Contacts {
			employeesRes[i].Contacts[ci] = model.ContactResponse{
				Label: cv.Label,
				Value: cv.Value,
				Type:  cv.Type,
			}
		}
	}

	return employeesRes, nil
}
