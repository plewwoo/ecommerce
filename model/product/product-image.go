package model

import "gorm.io/gorm"

type ProductImage struct {
	gorm.Model
	ProductID int    `json:"product_id" query:"product_id" gorm:"type:int(6)"`
	Image     string `json:"iamge" query:"iamge" gorm:"type:varchar(255)"`
	ImageUrl  string `json:"image_url" query:"image_url" gorm:"type:varchar(255)"`
}
