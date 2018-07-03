package list_controller

import (
	"github.com/astaxie/beego"
	"simple/models"
)

type IndexController struct {
	beego.Controller
}

func (this *IndexController) Prepare(){
	this.Layout = "layout.tpl"
}

func (this *IndexController) Get() {
	this.Data["list"] = models.GetAll()
	this.Layout = "layout.tpl"
	this.TplName = "list/index.tpl"
}