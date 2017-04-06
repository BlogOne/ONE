package controllers

import (
	"github.com/astaxie/beego"
	"ONE/models"
)

type ACatController struct {
	beego.Controller
}

func (this *ACatController) Get() {
	this.TplName = "AdminCategory.html"

	op := this.Input().Get("op")

	switch op {
	case "add":
		name := this.Input().Get("name")
		if len(name) == 0 {
			break
		}
		err := models.AddCategory(name)
		if err != nil {
			beego.Error(err)
		}

		this.Redirect("/admin/acategory", 301)

	case "del":
		id := this.Input().Get("id")
		if len(id) == 0 {
			break
		}

		err := models.DeleteCategory(id)
		if err != nil {
			beego.Error(err)
		}
		this.Redirect("/admin/acategory", 301)
		return

	}

	var err error
	this.Data["Categories"], err = models.GetAllCategories()

	if err != nil {
		beego.Error(err)
	}
}
