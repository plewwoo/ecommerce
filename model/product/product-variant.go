package model

import (
	"gorm.io/gorm"
)

type ProductVariant struct {
	gorm.Model
	ProductID     int    `json:"product_id" qurry:"product_id" gorm:"type:int(6)"`
	Description   string `json:"description" query:"description" gorm:"type:varchar(255)"`
	Price         int    `json:"price" query:"price" gorm:"type:int(10)"`
	CategoryID    int    `json:"category_id" query:"category_id" gorm:"type:int(6)"`
	SubCategoryID int    `json:"sub_category_id" query:"sub_category_id" gorm:"type:int(6)"`
	ColorID       int    `json:"color_id" query:"color_id" gorm:"type:int(6)"`
	SizeID        int    `json:"size_id" query:"size_id" gorm:"type:int(6)"`
	QuantityID    int    `json:"quantity_id" query:"quantity_id" gorm:"type:int(6)"`
}
