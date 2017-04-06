package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

type LoginController struct {
	beego.Controller
}

func (this *LoginController) Get() {
	isExit := this.Input().Get("exit") == "true"
	if isExit {
		this.Ctx.SetCookie("uname", "", -1, "/")
		this.Ctx.SetCookie("pwd", "", -1, "/")
		this.Redirect("/", 302)
		return
	}

	this.TplName = "login.html"
}

func (this *LoginController) Post() {
	// this.Ctx.WriteString(fmt.Sprint(this.Input())) //调用fmt,直接打印，忽略类型不一样
	// return
	uname := this.Input().Get("uname")
	pwd := this.Input().Get("pwd")
	autologin := this.Input().Get("autologin") == "on"

	if beego.AppConfig.String("uname") == uname &&
		beego.AppConfig.String("pwd") == pwd {
		maxAge := 0
		if autologin {
			maxAge = 3600 * 24
		}

		this.Ctx.SetCookie("uname", uname, maxAge, "/")
		this.Ctx.SetCookie("pwd", pwd, maxAge, "/")
		this.Redirect("/", 302)
		return
	} else {
		this.Abort("404")
		return
	}
	this.Redirect("/", 302)
	return
}

func checkAccount(ctx *context.Context) bool {
	ck, err := ctx.Request.Cookie("uname")
	if err != nil {
		return false
	}

	uname := ck.Value

	ck, err = ctx.Request.Cookie("pwd")
	if err != nil {
		return false
	}

	pwd := ck.Value
	return uname == beego.AppConfig.String("uname") &&
		pwd == beego.AppConfig.String("pwd")
}
