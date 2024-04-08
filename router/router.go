package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"

	authRoutes "ecommerce/routes/auth/auth"
	productRoutes "ecommerce/routes/product/product"
	productCategoryRoutes "ecommerce/routes/product/product-category"
	productColorRoutes "ecommerce/routes/product/product-color"
	productQuantityRoutes "ecommerce/routes/product/product-quantity"
	productSizeRoutes "ecommerce/routes/product/product-size"
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
		return ctx.SendString("Welcome Fiber")
	})

	api := app.Group("/api/v1")
	authRoutes.Setup(api)
	productRoutes.Setup(api)
	productVariantRoutes.Setup(api)
	productColorRoutes.Setup(api)
	productSizeRoutes.Setup(api)
	productQuantityRoutes.Setup(api)
	productCategoryRoutes.Setup(api)
	productSubCategoryRoutes.Setup(api)

	return app
}
