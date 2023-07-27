package app

var AppConfigTemplate = `package config

import (
	"{{ .moduleName }}/app/facades"
	"{{ .moduleName }}/app/providers"
)

type Providers struct {
	*providers.ControllerServiceProvider
}

type Facades struct {
	*facades.ControllerFacade
}
`