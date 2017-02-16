package controllers

import (
	"github.com/astaxie/beego"
)

type BaseController struct {
	beego.Controller
	Userid   int
	Username string
}

func (this *BaseController) Prepare() {
	loginId := this.GetSession("UserId")
	loginName := this.GetSession("UserName")
	if loginId != nil {
		this.Userid = loginId.(int)
		this.Username = loginName.(string)
	} else {
		this.Userid = 0
		this.Username = ""
	}
}

//在前端显示错误信息
func (this *BaseController) Showmsg(msg ...string) {
	if len(msg) == 1 {
		msg = append(msg, this.Ctx.Request.Referer())
	}
	this.Data["adminid"] = this.Userid
	this.Data["adminname"] = this.Username
	this.Data["msg"] = msg[0]
	this.Data["redirect"] = msg[1]
	this.Layout = "admin/layout.html"
	this.TplName = "admin/showmsg.html"
	this.Render()
	this.StopRun()
}
