package app

var CtrlFacadeTemplate = `
package facades

import (
	"{{ .moduleName }}/app/concretes"
	"github.com/owenzhou/ginrbac/support/facades"
)

var C *concretes.Controllers

type ControllerFacade struct {
	*facades.Facade
}

func (c *ControllerFacade) GetFacadeAccessor() {
	C = c.App.Make("controller").(*concretes.Controllers)
}
`