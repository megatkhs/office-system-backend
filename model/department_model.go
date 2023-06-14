package model

import (
	"time"

	"gorm.io/gorm"
)

type Department struct {
	gorm.Model
	Name        string     `gorm:"comment:部署名"`
	Description string     `gorm:"comment:部署説明"`
	Employees   []Employee `gorm:"comment:所属従業員"`
}

type DepartmentResponse struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
