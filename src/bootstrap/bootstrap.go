package bootstrap

import (
	"fiber-boilerplate/src/server"
	"fiber-boilerplate/src/server/controllers"
	"fiber-boilerplate/src/services"
	"log"

	"go.uber.org/dig"
)

type Dependency struct {
	Constructor interface{}
	Interface   interface{}
	Token       string
}

func Load() {
	container := dig.New()
	deps := []Dependency{
		{
			Constructor: services.NewFooService,
			Interface:   new(services.IFooService),
			Token:       "FooService",
		},
		{
			Constructor: controllers.NewFooController,
			Interface:   new(controllers.IFooController),
			Token:       "FooController",
		},
	}

	for _, dep := range deps {
		err := container.Provide(
			dep.Constructor,
			dig.As(dep.Interface),
			dig.Name(dep.Token),
		)

		if err != nil {
			log.Fatalf(err.Error())
		}

	}

	container.
		val := container.Invoke(server.NewServer)

	if val != nil {
		log.Fatalf(val.Error())
	}

}
