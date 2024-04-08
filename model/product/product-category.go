package model

import (
	"gorm.io/gorm"
)

type ProductCategory struct {
	gorm.Model
	Name        string `json:"name" query:"name" gorm:"type:varchar(255)"`
	Description string `json:"description" query:"description" gorm:"type:varchar(255)"`
	Image       string `json:"image" query:"image" gorm:"type:varchar(255)"`
	ImageUrl    string `json:"image_url" query:"image_url" gorm:"type:varchar(255)"`
}
