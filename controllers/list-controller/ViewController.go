package list_controller

import (
	"github.com/astaxie/beego"
	"strconv"
	"simple/models"
)

type ViewController struct {
	beego.Controller
}

func (this *ViewController)  Get(){
	id,_ := strconv.Atoi(this.Ctx.Input.Param(":id"))
	this.Data["list"] = models.GetList(id)
	this.Layout = "layout.tpl"
	this.TplName = "list/detail.tpl"
}