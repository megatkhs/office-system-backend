package repository

import (
	"backend/model"

	"gorm.io/gorm"
)

type DepartmentRepository interface {
	GetAllDepartments(departments *[]model.Department) error
	GetDepartmentById(department *model.Department, id uint) error
}

type departmentRepository struct {
	db *gorm.DB
}

func NewDepartmentRepository(db *gorm.DB) DepartmentRepository {
	return &departmentRepository{db}
}

func (dr *departmentRepository) GetAllDepartments(departments *[]model.Department) error {
	err := dr.db.Model(&model.Department{}).Find(departments).Error
	return err
}

func (dr *departmentRepository) GetDepartmentById(department *model.Department, id uint) error {
	err := dr.db.Preload("Employees").First(department, id).Error
	return err
}
