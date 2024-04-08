package Routes

import (
	handler "ecommerce/handler/auth/auth"

	"github.com/gofiber/fiber/v2"
)

func Setup(router fiber.Router) {
	newRouter := router.Group("/auth")
	newRouter.Post("/register", handler.Register)
}
