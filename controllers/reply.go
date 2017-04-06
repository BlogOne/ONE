package controllers

import (
	"ONE/models"

	"github.com/astaxie/beego"
)

type ReplyController struct {
	beego.Controller
}

func (this *ReplyController) Add() {
	// this.TplName = "topic_view.html"

	tid := this.Input().Get("tid")
	err := models.AddReply(
		tid,
		this.Input().Get("nickname"),
		this.Input().Get("email"),
		this.Input().Get("comment"),
	)
	if err != nil {
		beego.Error(err)
	}

	this.Redirect("/topic/view/"+tid, 302)

}

func (this *ReplyController) Delete() {
	if !checkAccount(this.Ctx) {
		return
	}
	tid := this.Input().Get("tid")

	err := models.DeleteReply(this.Input().Get("rid"))
	if err != nil {
		beego.Error(err)
	}

	this.Redirect("/topic/view/"+tid, 302)
}
