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
		Size string `json:"size"`
	}

	payload := new(Payload)
	productSizeOrm := new(model.ProductSize)
	err := c.BodyParser(payload)
	if err != nil {
		return common.FiberReviewPayload(c)
	}

	productID, _ := strconv.Atoi(id)

	productSizeOrm.ID = uint(productID)
	productSizeOrm.Size = payload.Size

	errSave := common.Database.Model(productSizeOrm).Where("id = ?", productSizeOrm.ID).Updates(&productSizeOrm)

	if errSave.Error != nil {
		common.PrintError(" save error error ", errSave.Error.Error())
		common.FiberError(c, fiber.StatusBadRequest, "can't save")
	}

	return common.FiberSuccess(c)
}
