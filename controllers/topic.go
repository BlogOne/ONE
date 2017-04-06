package controllers

import (
	"ONE/models"

	"net/http"

	"github.com/astaxie/beego"
	"os"
	"strconv"
	"time"
)

type TopicController struct {
	beego.Controller
	http.Request
}

func (this *TopicController) Get() {
	this.Data["IsLogin"] = checkAccount(this.Ctx)
	this.Data["IsTopic"] = true
	this.TplName = "topic.html"

	topics, err := models.GetAllTopics("", false, true)

	if err != nil {
		beego.Error(err)
	} else {
		this.Data["Topics"] = topics
	}

}

func (this *TopicController) Post() {
	if !checkAccount(this.Ctx) {
		this.Redirect("/login", 302)
		return
	}

	// 解析表单
	tid := this.Input().Get("tid")
	title := this.Input().Get("title")
	content := this.Input().Get("content")
	// content := this.Request.FormValue("editormd-html-code")
	category := this.Input().Get("category")
	label := this.Input().Get("label")
	summary := this.Input().Get("summary")

	var err error
	if len(tid) == 0 {
		err = models.AddTopic(title, category, label, summary, content)
	} else {
		err = models.ModifyTopic(tid, title, category, summary, content, label)
	}

	if err != nil {
		beego.Error(err)
	}

	// this.Redirect("/topic", 302)
	this.Redirect("/topic/view/"+tid, 302)
}

func (this *TopicController) Add() {
	this.Data["IsLogin"] = checkAccount(this.Ctx)
	this.TplName = "topic_add.html"
	if !checkAccount(this.Ctx) {
		this.Redirect("/login", 302)
		return
	}

	categories, err := models.GetAllCategories()
	if err != nil {
		beego.Error(err)
	}
	this.Data["Categories"] = categories
}

func (this *TopicController) View() {
	this.TplName = "topic_view.html"

	topic, err := models.GetTopic(this.Ctx.Input.Params()["0"])
	if err != nil {
		beego.Error(err)
		this.Redirect("/", 302)
		return
	}

	this.Data["Topic"] = topic
	this.Data["Tid"] = this.Ctx.Input.Params()["0"]

	replies, err := models.GetAllReplies(this.Ctx.Input.Params()["0"])
	if err != nil {
		beego.Error(err)
		return
	}

	this.Data["Replies"] = replies
	this.Data["IsLogin"] = checkAccount(this.Ctx)
	this.Data["IsLogin2"] = checkAccount(this.Ctx)

}

func (this *TopicController) Modify() {
	this.TplName = "topic_modify.html"

	tid := this.Input().Get("tid")
	topic, err := models.GetTopic(tid)
	if err != nil {
		beego.Error(err)
		this.Redirect("/", 302)
		return
	}

	this.Data["Topic"] = topic
	this.Data["Tid"] = tid

	categories, err := models.GetAllCategories()
	if err != nil {
		beego.Error(err)
	}
	this.Data["Categories"] = categories
}

func (this *TopicController) Delete() {
	this.TplName = "topic.html"
	if !checkAccount(this.Ctx) {
		this.Redirect("/login", 302)
		return
	}

	err := models.DeleteTopic(this.Input().Get("tid"))
	// err := models.DeleteTopic(this.Ctx.Input.Params()["0"])
	if err != nil {
		beego.Error(err)
	}
	this.Redirect("/topic", 302)
	return
}

// 接收文件
func (this *TopicController) Upload() {
	// 获取本月日期
	now := time.Now().Format("2006/01")
	// 设置保存目录
	mpath := "upload/image/" + now + "/"
	// 创建目录
	os.MkdirAll(mpath, 0755)

	_, h, err := this.GetFile("editormd-image-file")
	if err != nil {
		this.Data["json"] = map[string]interface{}{"success": 0, "message": err.Error()}
		this.ServeJSON()
	}

	fpath := mpath + h.Filename

	for i := 0; ; i++ {
		_, err = os.Stat(fpath)
		if err == nil {
			fpath = mpath + strconv.Itoa(i) + h.Filename
		} else {
			break
		}
	}

	err = this.SaveToFile("editormd-image-file", fpath)
	if err != nil {
		this.Data["json"] = map[string]interface{}{"success": 0, "message": err.Error()}
	} else {
		this.Data["json"] = map[string]interface{}{"success": 1, "message": "文件上传成功！", "url": beego.AppConfig.String("url") + fpath[0:]}
	}

	this.ServeJSON()
}
