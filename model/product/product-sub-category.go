package model

import "gorm.io/gorm"

type ProductSubCategory struct {
	gorm.Model
	CategoryID  int    `json:"category_id" query:"category_id" gorm:"type:int(6)"`
	Name        string `json:"name" query:"name" gorm:"type:varchar(255)"`
	Description string `json:"description" query:"description" gorm:"type:varchar(255)"`
	Image       string `json:"image" query:"image" gorm:"type:varchar(255)"`
	ImageUrl    string `json:"image_url" query:"image_url" gorm:"type:varchar(255)"`
}
