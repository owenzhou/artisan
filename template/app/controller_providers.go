package app

var CtrlProviderTemplate = `
package providers

import (
	"{{ .moduleName }}/app/concretes"
	"github.com/owenzhou/ginrbac/contracts"
	"github.com/owenzhou/ginrbac/support"
)

type ControllerServiceProvider struct {
	*support.ServiceProvider
}

func (c *ControllerServiceProvider) Register() {
	c.App.Singleton("controller", func(app contracts.IApplication) interface{} {
		return app.InjectApp(concretes.NewController())
	})
}
`
