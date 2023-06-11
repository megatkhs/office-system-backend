package model

import (
	"gorm.io/gorm"
)

type Employee struct {
	gorm.Model
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Birthday     string `json:"birthday"`
	DepartmentID uint   `json:"department_id"`
	Contacts     []Contact
}
