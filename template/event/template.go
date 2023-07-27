package event

var Template = `//代码生成器生成
package {{.packageName}}

import (
	"github.com/owenzhou/ginrbac/utils"
)

//每个事件有多个监听者
type {{.eventName}} struct {
	//存放的数据
	data interface{}
	//监听列表
}

func (e *{{.eventName}}) Data() map[string]interface{} {
	return utils.Struct2Map(e.data)
}

func New{{.eventName}}(data interface{}) *{{.eventName}} {
	return &{{.eventName}}{data: data}
}
`
