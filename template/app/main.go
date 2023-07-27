package app

var MainTemplate = `
package main

import (
	"embed"
	"{{ .moduleName }}/config"
	"github.com/owenzhou/ginrbac/app"
	"{{ .moduleName }}/routes"
)

//go:embed views
var views embed.FS

func main() {
	app := app.NewApp(views)

	//注册自定义的服务及门面
	app.Register(new(config.Providers))
	app.Register(new(config.Facades))

	//注册web路由
	routes.Web(app)

	app.Run()
}

`
