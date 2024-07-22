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
	fooController controllers.IFooController
}

type ServerDependencies struct {
	dig.In
	FooController controllers.IFooController `name:"FooController"`
}

func NewServer(deps ServerDependencies) {
	app := Server{
		fooController: deps.FooController,
	}

	envs := config.NewEnv()

	api := fiber.New()

	// signal channel to capture system calls
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)

	// start shutdown goroutine
	go func() {
		// capture sigterm and other system call here
		<-sigCh
		log.Info("Shutting down server...")
		_ = api.Shutdown()
	}()

	routes.GeneralRoute(api)
	routes.SwaggerRoute(api)
	routes.FooRoute(api, app.fooController)
	routes.NotFoundRoute(api)

	api.Listen(fmt.Sprintf(":%d", envs.Port))

}
