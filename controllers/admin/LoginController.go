package admin

import (
	"github.com/astaxie/beego"
	"fmt"
)

type LoginController struct {
	beego.Controller
}

func (l *LoginController) Get() {
	fmt.Println("1111111111111111111111111")
	l.TplName = "admin/login.tpl"
}

type LoginafterController struct {
	beego.Controller
}

//func (this *LoginafterController) Post() {
//	a := this.GetString("username")
//	b := this.GetString("password")
//	f:=this.Ctx.Request.URL.Path//获取更多的信息可以通过上下文来获取
//	fmt.Println(f)
//	fmt.Println(a, b)
//}

type user struct {
	username string `form:"username"`
	password string `form:"password"`
}

func (this *LoginafterController) Post() {
	u := user{}
	err := this.ParseForm(&u)

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("1111")
	fmt.Println(u.username == "", u.password == "")
}
