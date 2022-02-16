package listener

var Template = `//代码生成器生成
package {{.packageName}}

import (
	"fmt"
	"github.com/owenzhou/ginrbac/contracts"
)

type {{.listenerName}} struct {
}

func (a *{{.listenerName}}) Handle(event contracts.IEvent) bool {
	fmt.Println("data:", event.Data())
	return true
}
`
