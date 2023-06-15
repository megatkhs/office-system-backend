package usecase

import (
	"backend/model"
	"backend/repository"
)

type DepartmentUsecase interface {
	GetAllDepartments() ([]model.DepartmentResponse, error)
	GetDepartmentById(id uint) (model.DepartmentDetailResponse, error)
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

func (du *departmentUsecase) GetDepartmentById(id uint) (model.DepartmentDetailResponse, error) {
	department := model.Department{}

	if err := du.dr.GetDepartmentById(&department, id); err != nil {
		return model.DepartmentDetailResponse{}, err
	}

	departmentRes := model.DepartmentDetailResponse{
		ID:          department.ID,
		Name:        department.Name,
		Description: department.Description,
		CreatedAt:   department.CreatedAt,
		UpdatedAt:   department.UpdatedAt,
		Employees:   make([]model.EmployeeResponse, len(department.Employees)),
	}
	for ei, ev := range department.Employees {
		departmentRes.Employees[ei] = model.EmployeeResponse{
			ID:           ev.ID,
			FirstName:    ev.FirstName,
			LastName:     ev.LastName,
			Birthday:     ev.Birthday,
			DepartmentID: ev.DepartmentID,
		}
	}

	return departmentRes, nil
}
