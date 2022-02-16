package controller

var Template = `//代码生成器生成
package {{.packageName}}

import (
	"github.com/owenzhou/ginrbac/app"
	. "github.com/owenzhou/ginrbac/support/facades"
)

type {{.controllerName}} struct {
	*app.Controller
}

func (ctrl *{{.controllerName}}) Index(c *app.Context){
	Log.WithFields(Fields{
		"title": "hello world",
	}).Info("visit home index")

	c.JSON(200, app.H{
		"title":  "hello world",
	})
}
`
var ResourceTemplate = `//代码生成器生成
package {{.packageName}}

import (
	"github.com/owenzhou/ginrbac/app"
	. "github.com/owenzhou/ginrbac/support/facades"
)

type {{.controllerName}} struct {
	*app.Controller
}

//列表
func (ctrl *{{.controllerName}}) Index(c *app.Context){
	Log.WithFields(Fields{
		"title": "hello world",
	}).Info("visit home index")

	c.JSON(200, app.H{
		"title":  "列表方法",
	})
}

//显示
func (ctrl *{{.controllerName}}) Show(c *app.Context){
	c.JSON(200, app.H{
		"title":  "显示方法",
	})
}

//创建
func (ctrl *{{.controllerName}}) Create(c *app.Context){
	c.JSON(200, app.H{
		"title":  "创建方法",
	})
}

//保存
func (ctrl *{{.controllerName}}) Store(c *app.Context){
	c.JSON(200, app.H{
		"title":  "保存方法",
	})
}

//编辑
func (ctrl *{{.controllerName}}) Edit(c *app.Context){
	c.JSON(200, app.H{
		"title":  "编辑方法",
	})
}

//更新
func (ctrl *{{.controllerName}}) Update(c *app.Context){
	c.JSON(200, app.H{
		"title":  "更新方法",
	})
}

//删除
func (ctrl *{{.controllerName}}) Destroy(c *app.Context){
	c.JSON(200, app.H{
		"title":  "删除方法",
	})
}
`
