package bootstrap

import (
	"fiber-boilerplate/src/bootstrap/modules"
	"fiber-boilerplate/src/config"
	"fiber-boilerplate/src/dataAccess"
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

	// Load configuration
	env := config.LoadEnv()

	// Define dependencies
	deps := []Dependency{
		{
			Constructor: func() *config.Env {
				return env
			},
			Name: "Env",
		},
		{
			Constructor: func(env *config.Env) (*modules.MongoDB, error) {
				return modules.NewMongoDB(env.MongoDbUrl)
			},
			Name: "MongoDB",
		},
		{
			Constructor: dataAccess.NewFooDataAccess,
			Interface:   new(dataAccess.IFooDataAccess),
			Name:        "FooDataAccess",
		},
		{
			Constructor: controllers.NewFooController,
			Interface:   new(controllers.IFooController),
			Name:        "FooController",
		},
		{
			Constructor: services.NewFooService,
			Interface:   new(services.IFooService),
			Name:        "FooService",
		},
	}

	for _, dep := range deps {
		if dep.Interface == nil {
			// Directly provide the struct without dig.As
			err := container.Provide(dep.Constructor)
			if err != nil {
				log.Fatalf("Failed to provide dependency %s: %v", dep.Name, err)
			}
		} else {
			// Provide with interface
			err := container.Provide(
				dep.Constructor,
				dig.As(dep.Interface),
				dig.Name(dep.Name),
			)
			if err != nil {
				log.Fatalf("Failed to provide dependency %s: %v", dep.Name, err)
			}
		}
	}

	err := container.Invoke(server.NewServer)
	if err != nil {
		log.Fatalf("Failed to invoke server: %v", err)
	}
}
