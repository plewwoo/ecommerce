package Routes

import (
	handler "ecommerce/handler/product/product-color"

	"github.com/gofiber/fiber/v2"
)

func Setup(router fiber.Router) {
	newRouter := router.Group("/product-color")
	newRouter.Get("/", handler.GetData)
	newRouter.Post("/add", handler.Add)
	newRouter.Put("/:id", handler.Update)
	newRouter.Delete("/:id", handler.Delete)
}
