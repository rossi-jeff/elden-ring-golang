package models

import "gorm.io/gorm"

type Material struct {
	gorm.Model
	Name        string `gorm:"size:100;index:material_name;"`
	Description string `gorm:"size:500"`
}
