package services

import (
	"fiber-boilerplate/src/dataAccess"

	"go.uber.org/dig"
)

type IFooService interface {
	GetFoo(id string) (string, error)
}

type FooService struct {
	fooDataAccess dataAccess.IFooDataAccess
}

type FooServiceDependencies struct {
	dig.In

	FooDataAccess dataAccess.IFooDataAccess `name:"FooDataAccess"`
}

func NewFooService(deps FooServiceDependencies) *FooService {
	return &FooService{
		fooDataAccess: deps.FooDataAccess,
	}
}

func (service *FooService) GetFoo(id string) (string, error) {
	return service.fooDataAccess.GetFoo(id)
}
