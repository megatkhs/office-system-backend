package model

import "gorm.io/gorm"

type Department struct {
	gorm.Model
	Name        string `gorm:"comment:部署名"`
	Description string `gorm:"comment:部署説明"`
}
