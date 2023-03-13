package controllers

import (
	"log"

	beego "github.com/beego/beego/v2/server/web"
)

type HomeController struct {
	beego.Controller
}

func (c *HomeController) Get() {
	c.Data["Website"] = c.Ctx.Input.Param(":ext")
	log.Println(c.Data["Website"])
	c.Data["Email"] = "smith@gmail.com"
	c.TplName = "home.tpl"
}
