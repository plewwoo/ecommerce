package Routes

import (
	handler "ecommerce/handler/auth/auth"

	"github.com/gofiber/fiber/v2"
)

func Setup(router fiber.Router) {
	registerRoute := router.Group("/register")
	registerRoute.Post("/", handler.Register)

	loginRoute := router.Group("/login")
	loginRoute.Post("/", handler.Login)
}
