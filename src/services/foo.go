package services

import "go.uber.org/dig"

type IFooService interface {
	GetFoo(id string) string
}

type FooService struct {
}

type FooServiceDependencies struct {
	dig.In
}

func NewFooService(deps FooServiceDependencies) *FooService {
	return &FooService{}
}

func (service *FooService) GetFoo(id string) string {
	return id
}
