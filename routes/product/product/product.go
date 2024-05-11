package Routes

import (
	handler "ecommerce/handler/product/product"

	"github.com/gofiber/fiber/v2"
)

func Setup(router fiber.Router) {
	newRouter := router.Group("/product")
	newRouter.Get("/", handler.GetData)
	newRouter.Get("/:id", handler.GetDataByID)
	newRouter.Post("/add", handler.Add)
	newRouter.Put("/:id", handler.Update)
	newRouter.Delete("/:id", handler.Delete)
}
