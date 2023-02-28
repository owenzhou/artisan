package app

var GoModTemplate = `
module {{ .module }}

go 1.18

require (
	github.com/owenzhou/ginrbac@latest
)
`
