package handler

import (
	"ecommerce/common"
	model "ecommerce/model/product"

	"github.com/gofiber/fiber/v2"
)

func Add(c *fiber.Ctx) error {

	type Payload struct {
		Description   string `json:"description"`
		Price         int    `json:"price"`
		CategoryID    int    `json:"category_id"`
		SubCategoryID int    `json:"sub_category_id"`
		ColorName     string `json:"color_name"`
		ColorCode     string `json:"color_code"`
		SizeName      string `json:"size_name"`
		Size          string `json:"size"`
		Quantity      int    `json:"quantity"`
	}

	payload := new(Payload)
	err := c.BodyParser(payload)
	if err != nil {
		return common.FiberReviewPayload(c)
	}

	productVariantOrm := new(model.ProductVariant)
	productVariantOrm.Description = payload.Description
	productVariantOrm.Price = payload.Price
	productVariantOrm.CategoryID = payload.CategoryID
	productVariantOrm.SubCategoryID = payload.SubCategoryID
	productVariantOrm.SizeName = payload.SizeName
	productVariantOrm.Size = payload.Size
	productVariantOrm.ColorName = payload.ColorName
	productVariantOrm.ColorCode = payload.ColorCode
	productVariantOrm.Quantity = payload.Quantity

	errSave := common.Database.Create(productVariantOrm)

	if errSave.Error != nil {
		common.PrintError(" save error error ", errSave.Error.Error())
		return common.FiberError(c, "400", "can't save")
	}

	return common.FiberSuccess(c)
}
