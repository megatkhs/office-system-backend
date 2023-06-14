package usecase

import (
	"backend/model"
	"backend/repository"
)

type DepartmentUsecase interface {
	GetAllDepartments() ([]model.DepartmentResponse, error)
}

type departmentUsecase struct {
	dr repository.DepartmentRepository
}

func NewDepartmentUsecase(dr repository.DepartmentRepository) DepartmentUsecase {
	return &departmentUsecase{dr}
}

func (du *departmentUsecase) GetAllDepartments() ([]model.DepartmentResponse, error) {
	departments := []model.Department{}

	if err := du.dr.GetAllDepartments(&departments); err != nil {
		return nil, err
	}

	departmentsRes := make([]model.DepartmentResponse, len(departments))
	for i, v := range departments {
		departmentsRes[i] = model.DepartmentResponse{
			ID:          v.ID,
			Name:        v.Name,
			Description: v.Description,
		}
	}

	return departmentsRes, nil
}
