package routes

import (
	"fiber-boilerplate/src/server/controllers"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

func GeneralRoute(app *fiber.App) {
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"time":   time.Now().UnixNano(),
			"status": "OK",
		})
	})
}

func SwaggerRoute(app *fiber.App) {
	app.Get("/swagger/*", swagger.HandlerDefault) // default
}

func NotFoundRoute(app *fiber.App) {
	app.Use(
		func(c *fiber.Ctx) error {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"msg": "sorry, endpoint is not found",
			})
		},
	)
}

func FooRoute(app *fiber.App, controller controllers.IFooController) {
	app.Get("/foo/:id", controller.GetFoo)
}
