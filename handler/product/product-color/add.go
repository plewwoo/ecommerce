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
		Color     string `json:"color"`
	}

	payload := new(Payload)
	err := c.BodyParser(payload)
	if err != nil {
		return common.FiberReviewPayload(c)
	}

	productColorOrm := new(model.ProductColor)
	productColorOrm.ProductID = payload.ProductID
	productColorOrm.Name = payload.Name
	productColorOrm.Color = payload.Color

	errSave := common.Database.Create(productColorOrm)

	if errSave.Error != nil {
		common.PrintError(" save error error ", errSave.Error.Error())
		common.FiberError(c, fiber.StatusBadRequest, "can't save")
	}

	return common.FiberSuccess(c)
}
