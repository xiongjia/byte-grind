package util

import (
	"fmt"

	"github.com/samber/do"
)

type (
	Service1 interface {
		SvcPrint()
	}

	service1Impl struct{}

	CarService struct {
		Engine Service1
	}
)

func (s *service1Impl) HealthCheck() error {
	return nil
}

func (s *service1Impl) SvcPrint() {
	fmt.Println("svc1")
}

func NewSvc1(i *do.Injector) (Service1, error) {
	return &service1Impl{}, nil
}

func (s *CarService) SvcInvoke() {
	s.Engine.SvcPrint()
}

func NewCarService(i *do.Injector) (*CarService, error) {
	engine, err := do.Invoke[Service1](i)
	if err != nil {
		return nil, err
	}
	car := CarService{Engine: engine}
	return &car, nil
}

func SamberDo() {
	injector := do.New()
	do.Provide(injector, NewCarService)
	do.Provide(injector, NewSvc1)

	car := do.MustInvoke[*CarService](injector)
	car.SvcInvoke()
}
