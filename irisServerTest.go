package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
)

func irisSvr() {
	app := iris.Default()
	// Method:   GET
	// Resource: http://localhost:8080/
	app.Handle("GET", "/", func(ctx iris.Context) {
		ctx.HTML("Hello world!")
	})

	// same as app.Handle("GET", "/ping", [...])
	// Method:   GET
	// Resource: http://localhost:8080/ping
	app.Get("/ping", func(ctx iris.Context) {
		ctx.WriteString("pong")
	})

	// Method:   GET
	// Resource: http://localhost:8080/hello
	app.Get("/hello", func(ctx iris.Context) {
		ctx.JSON(iris.Map{"message": "Hello iris web framework."})
	})

	// http://localhost:8080
	// http://localhost:8080/ping
	// http://localhost:8080/hello
	app.Run(iris.Addr(":8080"))
}

func autoTLS() {
	app := iris.New()
	app.Get("/", func(ctx context.Context) {
		ctx.Writef("Hello from SECURE SERVER")
	})
	//必须用在真实的服务器上
	// app.Run(iris.AutoTLS("localhost:443"))
}

func configIris() {
	app := iris.New()
	app.Get("/", func(ctx context.Context) {
		ctx.HTML("<b>Hello!</b>")
	})

	//三种服务器配置方式
	// app.Configure(iris.WithConfiguration(iris.Configuration{DisableStartupLog: false}))
	// app.Run(iris.Addr(":8080"), iris.WithConfiguration(iris.Configuration{
	// 	DisableInterruptHandler: false,
	// }))

	// tmlPath := "./configs/irisConfigByTml.tml"
	// tmlConfig := iris.TOML(tmlPath)
	// app.Run(iris.Addr(":8080"), iris.WithConfiguration(tmlConfig))

	ymlPath := "./configs/irisConfigByYml.yml"
	app.Run(iris.Addr(":8080"), iris.WithConfiguration(iris.YAML(ymlPath)))

}
