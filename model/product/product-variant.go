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
	ColorName     string `json:"color_name" query:"color_name" gorm:"type:varchar(255)"`
	ColorCode     string `json:"color_code" query:"color_code" gorm:"type:varchar(10)"`
	SizeName      string `json:"size_name" query:"size_name" gorm:"type:varchar(255)"`
	Size          string `json:"size" query:"size" gorm:"type:varchar(255)"`
	Quantity      int    `json:"quantity_id" query:"quantity_id" gorm:"type:int(10)"`
}
