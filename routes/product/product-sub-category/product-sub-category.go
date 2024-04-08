package Routes

import (
	handler "ecommerce/handler/product/product-sub-category"

	"github.com/gofiber/fiber/v2"
)

func Setup(router fiber.Router) {
	newRouter := router.Group("/product-sub-category")
	newRouter.Get("/", handler.GetData)
	newRouter.Post("/add", handler.Add)
	newRouter.Put("/:id", handler.Update)
	newRouter.Delete("/:id", handler.Delete)
}
