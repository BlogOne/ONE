package routers

import (
	"ONE/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/tholeaf", &controllers.TholeafController{})
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/category", &controllers.CategoryController{})
	beego.Router("/topic", &controllers.TopicController{})

	beego.AutoRouter(&controllers.TopicController{})
	beego.Router("/reply", &controllers.ReplyController{})
	beego.Router("/reply/add", &controllers.ReplyController{}, "post:Add")
	beego.Router("/reply/delete", &controllers.ReplyController{}, "get:Delete")

	beego.Router("/admin", &controllers.AdminController{})
	beego.Router("/admin/acategory", &controllers.ACatController{})
	beego.Router("/admin/atopic", &controllers.ATopicController{})
	beego.Router("/upload", &controllers.TopicController{}, "post:Upload")
}
