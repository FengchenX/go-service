package service

import (
	"github.com/kataras/iris"
	"router"
)

type Svc struct {
	iris.Application
	Router  *router.Router
}

func NewSvc() *Svc {
	return &Svc{}
}

func DefaultSvc() *Svc {
	svc := NewSvc()
	svc.Application = *iris.Default()

	svc.Use()
	svc.Router = router.NewRouter(&svc.Application)
	return svc
}
