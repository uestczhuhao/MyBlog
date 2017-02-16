package admin

import (
	. "MyBlog/controllers"
	"MyBlog/models"
	"MyBlog/util"
	"strings"
	"time"
)

type LoginController struct {
	BaseController
}

func (this *LoginController) LoginIndex() {
	if this.Userid > 0 {
		this.Redirect("/admin", 302)
	}
	this.TplName = "admin/login/login.html"
}

func (this *LoginController) LoginPost() {
	if this.Userid > 0 {
		this.Redirect("/admin", 302)
	}
	username := this.GetString("account")
	password := this.GetString("password")

	if username != "" && password != "" {
		user := new(models.User)
		user.UserName = username
		// if user.Read("UserName") != nil || user.Read("Password") != nil {
		if user.Read("UserName") != nil || user.Password != util.Md5([]byte(password)) {
			this.Data["errmsg"] = "帐号或密码错误"
		} else if user.Active == 0 {
			this.Data["errmsg"] = "该帐号未激活"
		} else {
			user.LoginCount += 1
			// user.LastIp = this.getClientIp()
			this.SetSession("UserId", user.Id)
			this.SetSession("UserName", user.UserName)
			user.LastLogin = time.Now()
			user.Update()

			this.Redirect("/admin", 302)
		}
	}
	this.TplName = "admin/login/login.html"
}

func (this *LoginController) Loginupdate() {
	user := models.User{Id: this.Userid}
	if err := user.Read(); err != nil {
		this.Showmsg(err.Error())
	}

	if this.Ctx.Request.Method == "POST" {
		errmsg := make(map[string]string)
		password := strings.TrimSpace(this.GetString("password"))
		newpassword := strings.TrimSpace(this.GetString("newpassword"))
		newpassword2 := strings.TrimSpace(this.GetString("newpassword2"))
		updated := false
		if newpassword != "" {
			if password == "" || util.Md5([]byte(password)) != user.Password {
				errmsg["password"] = "当前密码错误"
			} else if len(newpassword) < 6 {
				errmsg["newpassword"] = "密码长度不能少于6个字符"
			} else if newpassword != newpassword2 {
				errmsg["newpassword2"] = "两次输入的密码不一致"
			}
			if len(errmsg) == 0 {
				user.Password = util.Md5([]byte(newpassword))
				user.Update("password")
				updated = true
			}
		}
		this.Data["updated"] = updated
		this.Data["errmsg"] = errmsg
	}
	this.Data["user"] = user
	this.Layout = "admin/layout.html"
	this.TplName = "admin/login/update.html"
}

func (this *LoginController) Loginout() {
	this.DestroySession()
	this.Redirect("/login", 302)
}
