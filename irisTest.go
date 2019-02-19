package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
)

func irisSvr() {
	app := iris.New()

	app.Get("/", homePageHandler)
	app.Run(iris.Addr(":8080"))
}

func homePageHandler(ctx context.Context) {
	ctx.Text("welcome to iris!")
}
