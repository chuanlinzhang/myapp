package admin

import (
	"github.com/astaxie/beego"
	"fmt"
)

type UseController struct {
	beego.Controller
}

func (u *UseController) Get()  {
fmt.Println("666")
}