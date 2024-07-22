package controllers

import (
	"fiber-boilerplate/src/services"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/dig"
)

type IFooController interface {
	GetFoo(c *fiber.Ctx) error
}

type FooController struct {
	fooService services.IFooService
}

type FooControllerDependencies struct {
	dig.In

	FooService services.IFooService `name:"FooService"`
}

func NewFooController(deps FooControllerDependencies) *FooController {
	return &FooController{
		fooService: deps.FooService,
	}
}

func (controller *FooController) GetFoo(c *fiber.Ctx) error {
	id := c.Params("id")

	foo := controller.fooService.GetFoo(id)

	return c.JSON(fiber.Map{
		"success": true,
		"foo":     foo,
	})
}
