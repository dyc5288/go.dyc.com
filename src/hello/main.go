package main

import (
	"github.com/astaxie/beego"
	"hello/controllers"
	_ "hello/routers"
)

func main() {
	beego.Router("/", &controllers.MainController{})
	beego.Run()
}
