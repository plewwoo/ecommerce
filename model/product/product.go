package model

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name        string `json:"name" query:"name" gorm:"type:varchar(255)"`
	Description string `json:"description" query:"description" gorm:"type:varchar(255)"`
}
