package controllers

import beego "github.com/beego/beego/v2/server/web"

type ApiController struct {
	beego.Controller
}

func (c *ApiController) List() {
	c.Data["Website"] = "list"
	c.Data["Email"] = "list@gmail.com"
	c.TplName = "api.tpl"
}

func (c *ApiController) Create() {
	c.Data["Website"] = "create"
	c.Data["Email"] = "create@gmail.com"
	c.TplName = "api.tpl"
}

func (c *ApiController) Update() {
	c.Data["Website"] = "udpate"
	c.Data["Email"] = "udpate@gmail.com"
	c.TplName = "api.tpl"
}

func (c *ApiController) Delete() {
	c.Data["Website"] = "delete"
	c.Data["Email"] = "delete@gmail.com"
	c.TplName = "api.tpl"
}
