package model

import (
	"time"

	"gorm.io/gorm"
)

// 部署情報 テーブルモデル
type Department struct {
	gorm.Model
	Name        string     `gorm:"comment:部署名"`
	Description string     `gorm:"comment:部署説明"`
	Employees   []Employee `gorm:"comment:所属従業員"`
}

// 部署情報 一覧レスポンスデータ
type DepartmentResponse struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

// 部署情報 詳細レスポンスデータ
type DepartmentDetailResponse struct {
	ID          uint               `json:"id"`
	Name        string             `json:"name"`
	Description string             `json:"description"`
	CreatedAt   time.Time          `json:"created_at"`
	UpdatedAt   time.Time          `json:"updated_at"`
	Employees   []EmployeeResponse `json:"employees"`
}
