package model

import (
	"time"

	"gorm.io/gorm"
)

// 従業員 テーブルモデル
type Employee struct {
	gorm.Model
	LastName     string    `gorm:"comment:名字"`
	FirstName    string    `gorm:"comment:名前"`
	Birthday     string    `gorm:"comment:誕生日"`
	DepartmentID uint      `gorm:"comment:部署ID"`
	Contacts     []Contact `gorm:"comment:連絡先"`
}

// 従業員 一覧レスポンスデータ
type EmployeeResponse struct {
	ID           uint   `json:"id"`
	LastName     string `json:"last_name"`
	FirstName    string `json:"first_name"`
	Birthday     string `json:"birthday"`
	DepartmentID uint   `json:"department_id"`
}

// 従業員 詳細レスポンスデータ
type EmployeeDetailResponse struct {
	ID           uint              `json:"id"`
	LastName     string            `json:"last_name"`
	FirstName    string            `json:"first_name"`
	Birthday     string            `json:"birthday"`
	DepartmentID uint              `json:"department_id"`
	CreatedAt    time.Time         `json:"created_at"`
	UpdatedAt    time.Time         `json:"updated_at"`
	Contacts     []ContactResponse `json:"contacts"`
}
