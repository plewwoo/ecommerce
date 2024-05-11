package handler

import (
	"ecommerce/common"
	model "ecommerce/model/product"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func Update(c *fiber.Ctx) error {
	id := c.Params("id")

	type Payload struct {
		Name        string `json:"name"`
		Description string `json:"description"`
		Image       string `json:"image"`
		ImageUrl    string `json:"image_url"`
	}

	payload := new(Payload)
	productCategoryOrm := new(model.ProductCategory)
	err := c.BodyParser(payload)
	if err != nil {
		return common.FiberReviewPayload(c)
	}

	productCategoryID, _ := strconv.Atoi(id)

	productCategoryOrm.ID = uint(productCategoryID)
	productCategoryOrm.Description = payload.Description
	productCategoryOrm.Image = payload.Image
	productCategoryOrm.ImageUrl = payload.ImageUrl

	errSave := common.Database.Model(productCategoryOrm).Where("id = ?", productCategoryOrm.ID).Updates(&productCategoryOrm)

	if errSave.Error != nil {
		common.PrintError(" save error error ", errSave.Error.Error())
		common.FiberError(c, "400", "can't save")
	}

	return common.FiberSuccess(c)
}
