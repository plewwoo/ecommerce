package handler

import (
	"ecommerce/common"
	model "ecommerce/model/product"

	"github.com/gofiber/fiber/v2"
)

func Add(c *fiber.Ctx) error {

	type Payload struct {
		CategoryID  int    `json:"category_id"`
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

	productSubCategoryOrm := new(model.ProductSubCategory)
	productSubCategoryOrm.CategoryID = payload.CategoryID
	productSubCategoryOrm.Name = payload.Name
	productSubCategoryOrm.Description = payload.Description
	productSubCategoryOrm.Image = payload.Image
	productSubCategoryOrm.ImageUrl = payload.ImageUrl

	errSave := common.Database.Create(productSubCategoryOrm)

	if errSave.Error != nil {
		common.PrintError(" save error error ", errSave.Error.Error())
		return common.FiberError(c, "400", "can't save")
	}

	return common.FiberSuccess(c)
}
