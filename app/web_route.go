package app

var WebRouteTemplate = `
package routes

import (
	"github.com/owenzhou/ginrbac/app"
	. "{{ .moduleName }}/app/facades"
)

func Web(a *app.App) {
	r := a.Router
	r.Get("/", C.HomeController.Index, "首页")
}
`
