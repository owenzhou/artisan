package app

var CtrlConcreteTemplate = `
package concretes

import (
	"{{ .moduleName }}/app/http/controllers"
)

type Controllers struct {
	//前台控制器
	*controllers.HomeController
}

func NewController() *Controllers {
	return new(Controllers)
}
`