package app

var CtrlTemplate = `
package {{ .packageName }}

import (
	"ginrbac/app/models"
	"ginrbac/bootstrap/app"
	. "ginrbac/bootstrap/support/facades"
)

type HomeController struct {
	*app.Controller
}

//初始化
func (ctrl *HomeController) Init(c *app.Context) {
	c.Share("url", c.Request.URL.Path)
}

//首页
func (ctrl *HomeController) Index(c *app.Context) {
	c.HTML(200, "/fronend/home/index", app.H{})
}
`
