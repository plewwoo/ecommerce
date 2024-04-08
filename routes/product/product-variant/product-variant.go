package Routes

import (
	handler "ecommerce/handler/product/product-variant"

	"github.com/gofiber/fiber/v2"
)

func Setup(router fiber.Router) {
	newRouter := router.Group("/product-variant")
	newRouter.Get("/", handler.GetData)
	newRouter.Post("/add", handler.Add)
	newRouter.Put("/:id", handler.Update)
	newRouter.Delete("/:id", handler.Delete)
}
