package list_controller

import (
	"github.com/astaxie/beego"
	"strconv"
	"simple/models"
)

type DeleteController struct {
	beego.Controller
}

func (this *DeleteController) Get() {
	id, _ := strconv.Atoi(this.Ctx.Input.Param(":id"))
	models.DelList(id)
	this.Ctx.Redirect(302, "/list")
}