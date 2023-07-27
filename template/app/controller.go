package app

var CtrlTemplate = `
package {{ .packageName }}

import (
	"github.com/owenzhou/ginrbac/app"
	. "github.com/owenzhou/ginrbac/support/facades"
)

type HomeController struct {
	*app.Controller
}

//初始化
func (ctrl *HomeController) Init(c *app.Context) {
	Log.WithFields(Fields{
		"title": "hello world",
	}).Info("visit home index")
	c.Share("url", c.Request.URL.Path)
}

//首页
func (ctrl *HomeController) Index(c *app.Context) {
	c.HTML(200, "/home/index", app.H{
		"title": "Hello world",
	})
}
`
