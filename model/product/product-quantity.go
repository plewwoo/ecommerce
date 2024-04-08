package model

import (
	"gorm.io/gorm"
)

type ProductQuantity struct {
	gorm.Model
	ProductID int `json:"product_id" query:"product_id" gorm:"type:int(6)"`
	Quantity  int `json:"quantity" query:"quantity" gorm:"type:int(6)"`
}
