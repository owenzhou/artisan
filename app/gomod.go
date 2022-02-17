package app

var GoModTemplate = `
module {{ .module }}

go 1.17

require (
	github.com/owenzhou/ginrbac v0.0.2
)
`
