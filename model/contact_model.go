package model

import "gorm.io/gorm"

// 連絡先情報 テーブルモデル
type Contact struct {
	gorm.Model
	Label      string `gorm:"comment:連絡先ラベル"`
	Value      string `gorm:"comment:連絡先"`
	Type       string `gorm:"comment:連絡先タイプ"`
	EmployeeID uint   `gorm:"comment:連絡先を紐づける従業員ID"`
}

// 連絡先情報 レスポンスデータ
type ContactResponse struct {
	Label string `json:"label"`
	Value string `json:"value"`
	Type  string `json:"type"`
}
