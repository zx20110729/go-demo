package routers

import (
	"web/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.SetViewsPath("/Users/zhou/code/github/go-demo/src/web/views")
	beego.Router("/", &controllers.MainController{})
	beego.Router("/user", &controllers.UserController{})
}
