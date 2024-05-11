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
	}

	payload := new(Payload)
	productOrm := new(model.Product)
	err := c.BodyParser(payload)
	if err != nil {
		return common.FiberReviewPayload(c)
	}

	productID, _ := strconv.Atoi(id)

	productOrm.ID = uint(productID)
	productOrm.Description = payload.Description

	errSave := common.Database.Model(productOrm).Where("id = ?", productOrm.ID).Updates(&productOrm)

	if errSave.Error != nil {
		common.PrintError(" save error error ", errSave.Error.Error())
		common.FiberError(c, "400", "can't save")
	}

	return common.FiberSuccess(c)
}
