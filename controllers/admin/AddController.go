package admin

import (
	"github.com/astaxie/beego"
	//"fmt"


)

type AddController struct {
	beego.Controller
}

//func (a *AddController) Get() {
//
//}
//func (a *AddController) AddGet() {
//
//	fmt.Println("11111111111111111")
//
//}
func (this *AddController) Prepare()  {

}
func (this *AddController) Get()  {
	this.Data["content"]="value"
	this.Layout="admin/layout.html"
	this.TplName="admin/add.tpl"
}
