package app

var MainTemplate = `
package main

import (
	"embed"
	"ginrbac/bootstrap/app"
	"ginrbac/routes"
)

//go:embed views
var views embed.FS

func main() {
	app := app.NewApp(views)

	//注册web路由
	routes.Web(app)

	app.Run()
}

`
