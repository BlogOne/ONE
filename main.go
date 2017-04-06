package main

import (
	"ONE/controllers"
	"ONE/models"
	_ "ONE/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	// "os"
)

func init() {
	models.RegisterDB()
}

func main() {

	orm.Debug = true
	orm.RunSyncdb("default", false, true)
	beego.ErrorController(&controllers.ErrorController{})

	beego.SetStaticPath("/upload", "upload")

	beego.Run()

}
