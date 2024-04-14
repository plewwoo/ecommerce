package config

import (
	// appModel "ecommerce/model/app"
	productModel "ecommerce/model/product"
	userModel "ecommerce/model/user"

	"gorm.io/gorm"
)

func AutoMigrate(DB *gorm.DB) {
	// DB.AutoMigrate(&appModel.AppInfo{})
	DB.AutoMigrate(&userModel.User{})
	DB.AutoMigrate(&userModel.UserAddress{})
	DB.AutoMigrate(&productModel.Product{})
	DB.AutoMigrate(&productModel.ProductVariant{})
	DB.AutoMigrate(&productModel.ProductImage{})
	DB.AutoMigrate(&productModel.ProductCategory{})
	DB.AutoMigrate(&productModel.ProductSubCategory{})
}
