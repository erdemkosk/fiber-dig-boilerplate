package routes

import (
	"fiber-boilerplate/src/server/controllers"

	"github.com/gofiber/fiber/v2"
)

func FooRoute(app *fiber.App, controller controllers.IFooController) {
	app.Get("/foo/:id", controller.GetFoo)
}
