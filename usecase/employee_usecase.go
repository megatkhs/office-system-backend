package usecase

import (
	"backend/model"
	"backend/repository"
)

type EmployeeUsecase interface {
	GetAllEmployees() ([]model.EmployeeResponse, error)
	GetEmployeeById(id uint) (model.EmployeeDetailResponse, error)
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
		}
	}

	return employeesRes, nil
}

func (eu *employeeUsecase) GetEmployeeById(id uint) (model.EmployeeDetailResponse, error) {
	employee := model.Employee{}

	if err := eu.er.GetEmployeeById(&employee, id); err != nil {
		return model.EmployeeDetailResponse{}, err
	}

	employeeRes := model.EmployeeDetailResponse{
		ID:           employee.ID,
		FirstName:    employee.FirstName,
		LastName:     employee.LastName,
		Birthday:     employee.Birthday,
		DepartmentID: employee.DepartmentID,
		CreatedAt:    employee.CreatedAt,
		UpdatedAt:    employee.UpdatedAt,
		Contacts:     make([]model.ContactResponse, len(employee.Contacts)),
	}
	for ci, cv := range employee.Contacts {
		employeeRes.Contacts[ci] = model.ContactResponse{
			Label: cv.Label,
			Value: cv.Value,
			Type:  cv.Type,
		}
	}

	return employeeRes, nil
}
