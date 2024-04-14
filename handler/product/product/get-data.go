package handler

import (
	"ecommerce/common"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func GetData(c *fiber.Ctx) error {

	type ProductImage struct {
		ID       uint   `json:"id"`
		Image    string `json:"image"`
		ImageUrl string `json:"image_url"`
	}

	type ProductVariant struct {
		ID            uint           `json:"id"`
		Description   string         `json:"description"`
		Price         int            `json:"price"`
		CategoryID    int            `json:"category_id"`
		SubCategoryID int            `json:"sub_category_id"`
		ColorName     string         `json:"color_name"`
		ColorCode     string         `json:"color_code"`
		SizeName      string         `json:"size_name"`
		Size          string         `json:"size"`
		Quantity      int            `json:"quantity"`
		Images        []ProductImage `json:"image"`
	}

	type Product struct {
		ID            uint             `json:"id"`
		Name          string           `json:"name"`
		Description   string           `json:"description"`
		CategoryID    int              `json:"category_id"`
		SubCategoryID int              `json:"sub_category_id"`
		Image         string           `json:"image"`
		ImageUrl      string           `json:"image_url"`
		Varaints      []ProductVariant `json:"varaint"`
	}

	var productResults []Product
	common.Database.Raw(`SELECT * FROM products WHERE deleted_at IS NULL`).Scan(&productResults)
	vehicleBrandModelResult := make([]Product, 0)

	for _, element := range productResults {
		var productVatiantOrm []ProductVariant

		fmt.Println(element)

		sql := `SELECT * FROM product_variants WHERE product_id = ?`
		common.Database.Raw(sql, element.ID).Scan(&productVatiantOrm)

		vehicleBrandModelResult = append(vehicleBrandModelResult, Product{ID: element.ID, Name: element.Name, Varaints: productVatiantOrm})
	}

	return c.JSON(fiber.Map{"status": 1, "code": "0000", "message": "success", "data": fiber.Map{
		"auction_year": vehicleBrandModelResult,
	},
	})
}
