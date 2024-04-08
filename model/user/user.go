package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Usercode  string `json:"usercode" query:"usercode" gorm:"type:varchar(255)"`
	Username  string `json:"username" query:"username" gorm:"type:varchar(255)"`
	Password  string `json:"password" query:"password" gorm:"type:varchar(255)"`
	Firstname string `json:"firstname" query:"firstname" gorm:"type:varchar(255)"`
	Lastname  string `json:"lastname" query:"lastname" gorm:"type:varchar(255)"`
	Tel       string `json:"tel" query:"tel" gorm:"type:varchar(255)"`
}
