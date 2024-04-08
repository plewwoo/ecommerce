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
		Name  string `json:"name"`
		Color string `json:"color"`
	}

	payload := new(Payload)
	productColorOrm := new(model.ProductColor)
	err := c.BodyParser(payload)
	if err != nil {
		return common.FiberReviewPayload(c)
	}

	productID, _ := strconv.Atoi(id)

	productColorOrm.ID = uint(productID)
	productColorOrm.Name = payload.Name
	productColorOrm.Color = payload.Color

	errSave := common.Database.Model(productColorOrm).Where("id = ?", productColorOrm.ID).Updates(&productColorOrm)

	if errSave.Error != nil {
		common.PrintError(" save error error ", errSave.Error.Error())
		common.FiberError(c, fiber.StatusBadRequest, "can't save")
	}

	return common.FiberSuccess(c)
}
