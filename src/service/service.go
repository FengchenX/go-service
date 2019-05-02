package service

import (
	"github.com/iris-contrib/middleware/cors"
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

	svc.Use(cors.AllowAll())

	svc.Router = router.NewRouter(&svc.Application)
	return svc
}
