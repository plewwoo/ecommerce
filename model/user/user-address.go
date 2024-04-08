package model

import "gorm.io/gorm"

type UserAddress struct {
	gorm.Model
	UserID     int    `json:"user_id" query:"user_id" gorm:"type:int(6)"`
	Address    string `json:"address" query:"address" gorm:"type:varchar(255)"`
	City       string `json:"city" query:"city" gorm:"type:varchar(255)"`
	PostalCode string `json:"postal_code" query:"postal_code" gorm:"type:varchar(5)"`
	Country    string `json:"country" query:"country" gorm:"type:varchar(255)"`
	Tel        string `json:"tel" query:"tel" gorm:"type:varchar(10)"`
}
