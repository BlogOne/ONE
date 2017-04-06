package controllers

import (
	"ONE/models"
	"github.com/astaxie/beego"
	"path"
)

type TholeafController struct {
	beego.Controller
}

func (this *TholeafController) Get() {
	this.TplName = "tholeaf.html"
	this.Data["IsTholeaf"] = true
	this.Data["IsLogin"] = checkAccount(this.Ctx)

	topics, err := models.GetAllTopics("", true, true)

	if err != nil {
		beego.Error(err)
	} else {
		this.Data["Topics"] = topics
	}
}

func (this *TholeafController) Post() {
	this.TplName = "tholeaf.html"
	if !checkAccount(this.Ctx) {
		this.Redirect("/login", 302)
		return
	}

	// 解析轻博的表单
	// summary := this.GetString("content")
	summary := this.Input().Get("content")

	_, fh, err := this.GetFile("file")
	if err != nil {
		beego.Error(err)
	}

	var image string

	if fh != nil {
		image = fh.Filename
		beego.Info(image)
		err := this.SaveToFile("file", path.Join("upload/image", image))
		if err != nil {
			beego.Error(err)
		}
	}

	err = models.AddTholeaf(summary, image)
	if err != nil {
		beego.Error(err)
	}

	this.Redirect("/tholeaf", 302)
}
