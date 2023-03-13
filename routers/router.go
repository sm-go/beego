package routers

import (
	"beego/controllers"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/home", &controllers.HomeController{})
	beego.Router("/api/list", &controllers.ApiController{}, "*:List")
	beego.Router("/api/create", &controllers.ApiController{}, "post:Create")
	beego.Router("/api/update", &controllers.ApiController{}, "put:Update")
	beego.Router("/api/delete", &controllers.ApiController{}, "delete:Delete")

}
