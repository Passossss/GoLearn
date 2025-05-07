package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.vip"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}
func (c *MainController) GetPortGu() {
	c.Data["Website"] = "https://www.linkedin.com/in/gustavo-passosss/"
	c.Data["Email"] = "gusapas26@gmail.com"
	c.TplName = "index.tpl"
}
