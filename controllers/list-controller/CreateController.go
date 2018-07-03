package list_controller


import (
	"github.com/astaxie/beego"
	"simple/models"
	"time"
)

type CreateController struct {
	beego.Controller
}

func (this *CreateController) Get() {
	this.Layout = "layout.tpl"
	this.TplName = "list/create.tpl"
}

func (this *CreateController) Post(){
	input := this.Input()
	var list models.Lists
	list.Title = input.Get("title")
	list.Description = input.Get("description")
	list.Created = time.Now()
	models.SaveList(list)
	this.Ctx.Redirect(302,"/list")
}