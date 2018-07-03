package list_controller

import (
	"github.com/astaxie/beego"
	"strconv"
	"simple/models"
	"time"
)

type EditController struct {
	beego.Controller
}

func (this *EditController) Get() {
	id,_ := strconv.Atoi(this.Ctx.Input.Param(":id"))
	this.Data["list"] = models.GetList(id)
	this.Layout = "layout.tpl"
	this.TplName = "list/edit.tpl"
}

func (this *EditController) Post(){
	inputs := this.Input()
	var list models.Lists

	list.Id,_ = strconv.Atoi(inputs.Get("id"))
	list.Title = inputs.Get("title")
	list.Description = inputs.Get("description")
	list.Created = time.Now()
	models.SaveList(list)

	this.Ctx.Redirect(302,"/list")
}

