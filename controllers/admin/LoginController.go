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
	l.TplName = "admin/login.html"
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
	Username string `form:"username"` //*****注意这里字段首字母必须大写
	Password string `form:"password"`
}

func (this *LoginafterController) Post() {
	this.SetSession("uid", int(2))
	v := this.GetSession("asta")
	if v == nil {
		this.SetSession("asta", int(1)) //设置session
		this.Data["num"] = 0
	} else {
		this.SetSession("asta", v.(int)+1)
		this.Data["num"] = v.(int)
	}
	fmt.Println(this.GetSession("asta")) //获取session
	u := user{}

	err := this.ParseForm(&u)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("1111")
	fmt.Println(u)
	//按照json的格式输出
	//this.Data["json"] = &u
	//this.ServeJSON()
	//按照XML的格式输出
	this.Data["xml"]=&u
	this.ServeXML()
}
