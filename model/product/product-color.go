package model

import (
	"gorm.io/gorm"
)

type ProductColor struct {
	gorm.Model
	ProductID int    `json:"product_id" query:"product_id" gorm:"type:int(6)"`
	Name      string `json:"name" query:"name" gorm:"type:varchar(255)"`
	Color     string `json:"color" query:"color" gorm:"type:varchar(255)"`
}
