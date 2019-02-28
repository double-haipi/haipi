package main

import "github.com/kataras/iris"

func basicRouter() {
	app := iris.New()
	app.OnErrorCode(iris.StatusNotFound, notFoundHandler)

}

func notFoundHandler(ctx iris.Context) {
	ctx.HTML("custom route for 404 not found")
	ctx.JSON(map[string]interface{}{
		"404": "not found",
	})
}
