package routers

import (
	"myapp/controllers"
	"github.com/astaxie/beego"
	"myapp/controllers/admin"

	"fmt"
	"github.com/astaxie/beego/context"
)

//验证用户是否已经登录，应用于全部的请求 过滤器
var FilterUser = func(ctx *context.Context) {
	_, ok := ctx.Input.Session("uid").(int)
	fmt.Println(ok, "*****") //false表示还没有登录
	if !ok && ctx.Request.RequestURI != "/admin/login" && ctx.Request.RequestURI != "/admin/loginafter" {
		ctx.Redirect(302, "/admin/login") //重定向
	}
}

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.InsertFilter("/*", beego.BeforeRouter, FilterUser)
	beego.Router("/admin", &admin.UseController{})
	//beego.Router("/admin/add",&admin.AddController{})
	//beego.Router("/admin/add", &admin.AddController{}, "get:AddGet")
	beego.Router("/admin/login", &admin.LoginController{})
	beego.Router("/admin/loginafter", &admin.LoginafterController{})
}

//一个固定的路由，一个控制器，然后根据用户请求方法不同请求控制器中对应的方法
