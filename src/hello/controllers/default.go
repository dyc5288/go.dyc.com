package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
	this.Data["Website"] = "beego.me2"
	this.Data["Email"] = "yunchao@gmail.com"
	this.TplNames = "index.tpl"
}
