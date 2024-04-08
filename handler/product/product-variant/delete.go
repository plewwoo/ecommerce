package handler

import (
	"ecommerce/common"
	model "ecommerce/model/product"

	"github.com/gofiber/fiber/v2"
)

func Delete(c *fiber.Ctx) error {
	id := c.Params("id")

	if id == "" {
		return common.FiberReviewPayload(c)
	}

	common.Database.Where("id = ?", id).Delete(&model.ProductVariant{})

	return common.FiberSuccess(c)
}
