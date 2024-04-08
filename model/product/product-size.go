package model

import (
	"gorm.io/gorm"
)

type ProductSize struct {
	gorm.Model
	ProductID int    `json:"product_id" query:"product_id" gorm:"type:int(6)"`
	Name      string `json:"name" query:"name" gorm:"type:varchar(255)"`
	Size      string `json:"size" query:"size" gorm:"type:varchar(2)"`
}
