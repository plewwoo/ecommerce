package handler

import (
	"ecommerce/common"

	"github.com/gofiber/fiber/v2"
)

func GetData(c *fiber.Ctx) error {

	var args []interface{}

	sql := `SELECT * FROM product_categories WHERE deleted_at IS NULL`

	return common.FiberQuery(c, sql, args...)
}
