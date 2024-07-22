package server

import (
	"fiber-boilerplate/src/config"
	"fiber-boilerplate/src/server/controllers"
	"fiber-boilerplate/src/server/routes"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"go.uber.org/dig"
)

type Server struct {
	app *fiber.App

	fooController controllers.IFooController
}

type ServerDependencies struct {
	dig.In

	FooController controllers.IFooController `name:"FooController"`
}

func NewServer(deps ServerDependencies) {
	server := &Server{
		app: fiber.New(),

		fooController: deps.FooController,
	}

	envs := config.LoadEnv()

	bindRoutes(*server)

	gracefulShutdown(*&server.app)

	server.app.Listen(fmt.Sprintf(":%d", envs.Port))

}

func bindRoutes(server Server) {
	routes.HealthRoute(server.app)
	routes.SwaggerRoute(server.app)
	routes.FooRoute(server.app, server.fooController)
	routes.NotFoundRoute(server.app)
}

func gracefulShutdown(app *fiber.App) {
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)

	go func() {
		<-sigCh
		log.Info("Shutting down server...")
		_ = app.Shutdown()
	}()
}
