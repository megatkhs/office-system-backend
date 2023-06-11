package model

import "gorm.io/gorm"

type Contact struct {
	gorm.Model
	Label      string `json:"label"`
	Value      string `json:"value"`
	Type       string `json:"type"`
	EmployeeID uint   `json:"employee_id"`
}
