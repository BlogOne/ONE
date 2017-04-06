package controllers

import (
	"github.com/astaxie/beego"
	//"strconv"
	"ONE/models"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["IsHome"] = true
	c.TplName = "home.html"
	c.Data["IsLogin"] = checkAccount(c.Ctx)

	topics, err := models.GetAllTopics(c.Input().Get("cate"), true, false)
	if err != nil {
		beego.Error(err)
	} else {
		c.Data["Topics"] = topics
	}

	categories, err := models.GetAllCategories()
	if err != nil {
		beego.Error(err)
	}

	c.Data["Categories"] = categories

	// 模板使用
	// c.Data["Website"] = "golang.org"
	// c.Data["Email"] = "astaxie@gmail.com"

	// c.Data["Truecond"] = true
	// c.Data["Falsecond"] = false

	// type u struct {
	// 	Name string
	// 	Age  int
	// 	Sex  string
	// }

	// user := &u{
	// 	Name: "GoBlog",
	// 	Age:  9,
	// 	Sex:  "nil",
	// }

	// c.Data["user"] = user

	// nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}

	// c.Data["Nums"] = nums

	// c.Data["GO"] = "hello GO"

	// c.Data["html"] = "<div> Hello Beego</div>"

	// c.Data["Pipe"] = "<div>Hello Beego</div>"

}
