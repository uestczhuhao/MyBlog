package admin

import (
	"MyBlog/util"

	"github.com/astaxie/beego"

	"MyBlog/models"
)

type LabelController struct {
	AdminIndexController
}

func (this *LabelController) Label() {
	var page int
	var pagesize int = 10
	var list []*models.Label
	var label models.Label

	if page, _ = this.GetInt("page"); page < 1 {
		page = 1
	}
	offset := (page - 1) * pagesize

	count, _ := label.Query().Count()
	if count > 0 {
		label.Query().OrderBy("-count").Limit(pagesize, offset).All(&list)
	}

	this.Data["count"] = count
	this.Data["list"] = list
	this.Data["pagebar"] = util.NewPager(page, int(count), pagesize, "/admin/label", true).ToString()
	this.Layout = "admin/layout.html"
	this.TplName = "admin/label/list.html"
}

func (this *LabelController) Add() {
	classify := this.GetString("newclass")
	var newLabel models.Label
	newLabel.Name = classify
	newLabel.Count = 0
	err := newLabel.Add()
	if err != nil {
		beego.Debug(err)
	}
	this.Redirect("/admin/label", 302)
}
