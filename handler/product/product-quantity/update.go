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
		ProductID int `json:"product_id"`
		Quantity  int `json:"quantity"`
	}

	payload := new(Payload)
	productQuantityOrm := new(model.ProductQuantity)
	err := c.BodyParser(payload)
	if err != nil {
		return common.FiberReviewPayload(c)
	}

	productID, _ := strconv.Atoi(id)

	productQuantityOrm.ID = uint(productID)
	productQuantityOrm.ProductID = payload.ProductID
	productQuantityOrm.Quantity = payload.Quantity

	errSave := common.Database.Model(productQuantityOrm).Where("id = ?", productQuantityOrm.ID).Updates(&productQuantityOrm)

	if errSave.Error != nil {
		common.PrintError(" save error error ", errSave.Error.Error())
		common.FiberError(c, fiber.StatusBadRequest, "can't save")
	}

	return common.FiberSuccess(c)
}
