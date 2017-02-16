package admin

import (
	. "MyBlog/controllers"
	"MyBlog/models"

	"os"
	"runtime"
)

type AdminIndexController struct {
	BaseController
}

func (this *AdminIndexController) Prepare() {
	this.CheckLogin()
}

//验证登录状态是否成功，不成功则跳回登录页面
func (this *AdminIndexController) CheckLogin() int {
	loginId := this.GetSession("UserId")
	loginName := this.GetSession("UserName")

	var loginIdInt int
	if loginId != nil {
		loginIdInt = loginId.(int)
		this.Userid = loginId.(int)
		this.Username = loginName.(string)
	} else {
		loginIdInt = 0
		this.Redirect("/login", 302)
	}
	return loginIdInt
}

func (this *AdminIndexController) Get() {
	this.Data["adminname"] = this.Username
	this.TplName = "admin/index.html"
}

func (this *AdminIndexController) Post() {
	this.Data["adminname"] = this.Username
	this.TplName = "admin/index.html"
}

func (this *AdminIndexController) System() {
	this.Data["hostname"], _ = os.Hostname()
	this.Data["version"] = "1.0.0"
	this.Data["gover"] = runtime.Version()
	this.Data["os"] = runtime.GOOS
	this.Data["cpunum"] = runtime.NumCPU()
	this.Data["arch"] = runtime.GOARCH

	this.Data["postnum"], _ = new(models.Article).Query().Count()
	this.Data["labelnum"], _ = new(models.Label).Query().Count()
	this.Data["usernum"], _ = new(models.User).Query().Count()
	this.Layout = "admin/layout.html"
	this.TplName = "admin/system/system.html"
}
