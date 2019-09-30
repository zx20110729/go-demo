package controllers

import (
	"github.com/astaxie/beego"
)

type UserController struct {
	beego.Controller
}

func (this *UserController) Get() {
	outStruct := JsonOut{200, "成功"}
	this.Data["json"] = &outStruct
	this.ServeJSON()
}
