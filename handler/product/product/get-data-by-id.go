package handler

import (
	"ecommerce/common"

	"github.com/gofiber/fiber/v2"
)

func GetDataByID(c *fiber.Ctx) error {

	id := c.Params("id")

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
		Varaint       []ProductVariant `json:"varaint"`
	}

	var productResult Product
	common.Database.Raw(`SELECT * FROM products WHERE id = ? AND deleted_at IS NULL`, id).Scan(&productResult)

	var productVariantOrm []ProductVariant
	sql := `SELECT * FROM product_variants WHERE product_id = ?`
	common.Database.Raw(sql, productResult.ID).Scan(&productVariantOrm)
	productVaraintResults := make([]ProductVariant, 0)

	for _, variant := range productVariantOrm {
		var productVariantImgOrm []ProductImage
		sql := `SELECT * FROM product_images WHERE product_variant_id = ?`
		common.Database.Raw(sql, variant.ID).Scan(&productVariantImgOrm)
		productVaraintResults = append(productVaraintResults, ProductVariant{ID: variant.ID, Description: variant.Description, Price: variant.Price, CategoryID: variant.CategoryID, SubCategoryID: variant.SubCategoryID, ColorName: variant.ColorName, ColorCode: variant.ColorCode, SizeName: variant.SizeName, Size: variant.Size, Quantity: variant.Quantity, Images: productVariantImgOrm})
	}
	productResult.Varaint = productVaraintResults

	return c.JSON(fiber.Map{"status": 1, "code": "0000", "message": "success", "data": productResult})
}
