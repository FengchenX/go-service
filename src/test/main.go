package main

import (
	"github.com/kataras/iris"
)

func main() {
	app := iris.Default()
	env := app.Party("/env-finishing")
	env.Get("/compensate", func(c iris.Context) {
		c.JSON("monitor")
	})
	app.Get("test", func(c iris.Context) {
		c.JSON("test")
	})
	app.Run(iris.Addr(":8081"))
}
