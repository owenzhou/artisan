package app

var WebRouteTemplate = `
package routes

import (
	"ginrbac/bootstrap/app"
)

func Web(a *app.App) {
	r := a.Router
	r.Get("/", C.HomeController.Index, "首页")
}
`
