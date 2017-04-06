package controllers

import (
	"github.com/astaxie/beego"
)

type ErrorController struct {
	beego.Controller
}

func (this *ErrorController) Error404() {
	this.Data["content"] = "page not found"
	this.TplName = "404.html"
}

func (this *ErrorController) Error501() {
    this.Data["content"] = "server error"
    this.TplName = "500.html"
}

// func (this *ErrorController) ErrorDb() {
//     this.Data["content"] = "database is now down"
//     this.TplName = "dberror.tpl"
// }
