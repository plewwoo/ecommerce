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
		CategoryID    int    `json:"cateogory_id"`
		SubCategoryID int    `json:"sub_category_id"`
		SizeID        int    `json:"size_id"`
		ColorID       int    `json:"color_id"`
		QuantityID    int    `json:"quantity_id"`
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
	productVariantOrm.SizeID = payload.SizeID
	productVariantOrm.ColorID = payload.ColorID
	productVariantOrm.QuantityID = payload.QuantityID

	errSave := common.Database.Create(productVariantOrm)

	if errSave.Error != nil {
		common.PrintError(" save error error ", errSave.Error.Error())
		common.FiberError(c, fiber.StatusBadRequest, "can't save")
	}

	return common.FiberSuccess(c)
}
