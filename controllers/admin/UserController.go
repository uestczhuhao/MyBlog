package admin

import (
	"MyBlog/models"
	"MyBlog/util"
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
)

type UserController struct {
	AdminIndexController
}

func (this *UserController) Add() {
	input := make(map[string]string)
	errmsg := make(map[string]string)
	if this.Ctx.Request.Method == "POST" {
		username := strings.TrimSpace(this.GetString("username"))
		password := strings.TrimSpace(this.GetString("password"))
		password2 := strings.TrimSpace(this.GetString("password2"))
		email := strings.TrimSpace(this.GetString("email"))
		active, _ := this.GetInt("active")

		input["username"] = username
		input["password"] = password
		input["password2"] = password2
		input["email"] = email

		valid := validation.Validation{}

		if v := valid.Required(username, "username"); !v.Ok {
			errmsg["username"] = "请输入用户名"
		} else if v := valid.MaxSize(username, 15, "username"); !v.Ok {
			errmsg["username"] = "用户名长度不能大于15个字符"
		}

		if v := valid.Required(password, "password"); !v.Ok {
			errmsg["password"] = "请输入密码"
		}

		if v := valid.Required(password2, "password2"); !v.Ok {
			errmsg["password2"] = "请再次输入密码"
		} else if password != password2 {
			errmsg["password2"] = "两次输入的密码不一致"
		}

		if v := valid.Required(email, "email"); !v.Ok {
			errmsg["email"] = "请输入email地址"
		} else if v := valid.Email(email, "email"); !v.Ok {
			errmsg["email"] = "Email无效"
		}
		beego.Debug(errmsg)
		if active > 0 {
			active = 1
		} else {
			active = 0
		}

		if len(errmsg) == 0 {
			var user models.User
			user.UserName = username
			user.Password = util.Md5([]byte(password))
			user.Email = email
			user.Active = int8(active)
			if err := user.Add(); err != nil {
				this.Showmsg(err.Error())
			}
			this.Redirect("/admin/user/list", 302)
		}

	}

	this.Data["input"] = input
	this.Data["errmsg"] = errmsg
	this.Layout = "admin/layout.html"
	this.TplName = "admin/user/add.html"
}

func (this *UserController) List() {
	var (
		page     int
		pagesize int = 10
		list     []*models.User
		user     models.User
	)

	if page, _ = this.GetInt("page"); page < 1 {
		page = 1
	}
	offset := (page - 1) * pagesize

	count, _ := user.Query().Count()
	if count > 0 {
		user.Query().OrderBy("-id").Limit(pagesize, offset).All(&list)
	}

	this.Data["count"] = count
	this.Data["list"] = list
	this.Data["pagebar"] = util.NewPager(page, int(count), pagesize, "/admin/user/list", true).ToString()
	this.Layout = "admin/layout.html"
	this.TplName = "admin/user/list.html"
}

func (this *UserController) Edit() {
	this.Layout = "admin/layout.html"
	this.TplName = "admin/user/edit.html"
}

func (this *UserController) Delete() {

}

func (this *UserController) Save() {

	input := make(map[string]string)
	errmsg := make(map[string]string)

	username := strings.TrimSpace(this.GetString("username"))
	password := strings.TrimSpace(this.GetString("password"))
	password2 := strings.TrimSpace(this.GetString("password2"))
	email := strings.TrimSpace(this.GetString("email"))
	active, _ := this.GetInt("active")

	input["username"] = username
	input["password"] = password
	input["password2"] = password2
	input["email"] = email

	valid := validation.Validation{}

	if v := valid.Required(username, "username"); !v.Ok {
		errmsg["username"] = "请输入用户名"
	} else if v := valid.MaxSize(username, 15, "username"); !v.Ok {
		errmsg["username"] = "用户名长度不能大于15个字符"
	}
	beego.Debug(errmsg)
	if v := valid.Required(password, "password"); !v.Ok {
		errmsg["password"] = "请输入密码"
	}

	if v := valid.Required(password2, "password2"); !v.Ok {
		errmsg["password2"] = "请再次输入密码"
	} else if password != password2 {
		errmsg["password2"] = "两次输入的密码不一致"
	}

	if v := valid.Required(email, "email"); !v.Ok {
		errmsg["email"] = "请输入email地址"
	} else if v := valid.Email(email, "email"); !v.Ok {
		errmsg["email"] = "Email无效"
	}
	beego.Debug(errmsg)

	if active > 0 {
		active = 1
	} else {
		active = 0
	}

	if len(errmsg) == 0 {
		var user models.User
		user.UserName = username
		user.Password = util.Md5([]byte(password))
		user.Email = email
		user.Active = int8(active)
		if err := user.Add(); err != nil {
			this.Showmsg(err.Error())
		}
		this.Redirect("/admin/user/list", 302)
	}

	this.Data["input"] = input
	this.Data["errmsg"] = errmsg
	this.Redirect("/admin/user/list", 302)
}
