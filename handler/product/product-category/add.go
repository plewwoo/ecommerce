package handler

import (
	"ecommerce/common"
	model "ecommerce/model/product"

	"github.com/gofiber/fiber/v2"
)

func Add(c *fiber.Ctx) error {

	type Payload struct {
		Name        string `json:"name"`
		Description string `json:"description"`
		Image       string `json:"image"`
		ImageUrl    string `json:"image_url"`
	}

	payload := new(Payload)
	err := c.BodyParser(payload)
	if err != nil {
		return common.FiberReviewPayload(c)
	}

	productCategoryOrm := new(model.ProductCategory)
	productCategoryOrm.Name = payload.Name
	productCategoryOrm.Description = payload.Description
	productCategoryOrm.Image = payload.Image
	productCategoryOrm.ImageUrl = payload.ImageUrl

	errSave := common.Database.Create(productCategoryOrm)

	if errSave.Error != nil {
		common.PrintError(" save error error ", errSave.Error.Error())
		common.FiberError(c, fiber.StatusBadRequest, "can't save")
	}

	return common.FiberSuccess(c)
}
