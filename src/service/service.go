package service

import (
	"github.com/iris-contrib/middleware/cors"
	"github.com/kataras/iris"
)

type Svc struct {
	iris.Application
}

func NewSvc() *Svc {
	return &Svc{}
}

func DefaultSvc() *Svc {
	svc := NewSvc()
	svc.Application = *iris.Default()

	svc.Use(cors.AllowAll())
	return svc
}
