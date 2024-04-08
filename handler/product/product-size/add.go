package handler

import (
	"ecommerce/common"
	model "ecommerce/model/product"

	"github.com/gofiber/fiber/v2"
)

func Add(c *fiber.Ctx) error {

	type Payload struct {
		ProductID int    `json:"product_id"`
		Name      string `json:"name"`
		Size      string `json:"size"`
	}

	payload := new(Payload)
	err := c.BodyParser(payload)
	if err != nil {
		return common.FiberReviewPayload(c)
	}

	productSizeOrm := new(model.ProductSize)
	productSizeOrm.ProductID = payload.ProductID
	productSizeOrm.Size = payload.Size

	errSave := common.Database.Create(productSizeOrm)

	if errSave.Error != nil {
		common.PrintError(" save error error ", errSave.Error.Error())
		common.FiberError(c, fiber.StatusBadRequest, "can't save")
	}

	return common.FiberSuccess(c)
}
