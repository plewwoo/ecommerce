package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"

	authRoutes "ecommerce/routes/auth/auth"
	productRoutes "ecommerce/routes/product/product"
	productCategoryRoutes "ecommerce/routes/product/product-category"
	productSubCategoryRoutes "ecommerce/routes/product/product-sub-category"
	productVariantRoutes "ecommerce/routes/product/product-variant"
)

func SetupRouter() *fiber.App {
	app := fiber.New()

	// Middleware
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "*",
		AllowHeaders: "*",
	}))
	app.Use(logger.New())

	// HTTP Method
	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.SendString("Ecommerce API")
	})

	api := app.Group("/api/v1")
	authRoutes.Setup(api)
	productRoutes.Setup(api)
	productVariantRoutes.Setup(api)
	productCategoryRoutes.Setup(api)
	productSubCategoryRoutes.Setup(api)

	return app
}
