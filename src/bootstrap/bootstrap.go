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
	Name        string
}

func Init() {
	container := dig.New()
	deps := []Dependency{
		{
			Constructor: services.NewFooService,
			Interface:   new(services.IFooService),
			Name:        "FooService",
		},
		{
			Constructor: controllers.NewFooController,
			Interface:   new(controllers.IFooController),
			Name:        "FooController",
		},
	}

	for _, dep := range deps {
		err := container.Provide(
			dep.Constructor,
			dig.As(dep.Interface),
			dig.Name(dep.Name),
		)

		if err != nil {
			log.Fatalf(err.Error())
		}
	}

	err := container.Invoke(server.NewServer)

	if err != nil {
		panic(err.Error())
	}

}
