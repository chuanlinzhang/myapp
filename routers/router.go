package routers

import (
	"myapp/controllers"
	"github.com/astaxie/beego"
	"myapp/controllers/admin"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/admin", &admin.UseController{})
	//beego.Router("/admin/add",&admin.AddController{})
	//beego.Router("/admin/add", &admin.AddController{}, "get:AddGet")
	beego.Router("/admin/login", &admin.LoginController{})
	beego.Router("/admin/loginafter", &admin.LoginafterController{})
}

//一个固定的路由，一个控制器，然后根据用户请求方法不同请求控制器中对应的方法
