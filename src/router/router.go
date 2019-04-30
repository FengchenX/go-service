package router

import (
	"github.com/kataras/iris"
)

type Router struct {
	App *iris.Application
}

func NewRouter(app *iris.Application) *Router {
	return &Router{App: app}
}
