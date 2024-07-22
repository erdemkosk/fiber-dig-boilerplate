package routes

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

func HealthRoute(app *fiber.App) {
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"time":   time.Now().UnixNano(),
			"status": "OK",
		})
	})
}
