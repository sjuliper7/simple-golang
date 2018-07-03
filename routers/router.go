package routers

import (
	"simple/controllers"
	"github.com/astaxie/beego"
	"simple/controllers/list-controller"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/list",&list_controller.IndexController{})
	beego.Router("/list-new",&list_controller.CreateController{})
	beego.Router("/list-edit/:id([0-9]+)", &list_controller.EditController{})
	beego.Router("/list-view/:id([0-9]+)", &list_controller.ViewController{})
	beego.Router("/list-delete/:id([0-9]+)", &list_controller.DeleteController{})
	//beego.Router("/llist-detail/:id([0-9]+)", &list_controller.)
}
