package models

import (
	"fmt"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type Article struct {
	Id         int
	UserId     int       `orm:"index"`
	Author     string    `orm:"size(32)"`
	Title      string    `orm:"size(128)"`
	Color      string    `orm:"size(7)"`
	Content    string    `orm:"type(text)"`
	PostTime   time.Time `orm:"type(datetime)"`
	Views      int
	Status     int8      ////0为已发布，1为草稿，2为回收站
	UpdateTime time.Time `orm:"type(datetime)"`
	IsTop      int8
	Labels     []*Label `orm:"reverse(many)"`
}

func (a *Article) Add() error {
	if _, err := orm.NewOrm().Insert(a); err != nil {
		return err
	}
	return nil
}

func (a *Article) Delete() error {
	if _, err := orm.NewOrm().Delete(a); err != nil {
		return err
	}
	return nil
}

func (a *Article) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(a, fields...); err != nil {
		return err
	}
	return nil
}

func (a *Article) Read(fields ...string) error {
	if err := orm.NewOrm().Read(a, fields...); err != nil {
		return err
	}
	return nil
}

func (a *Article) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(a)
}

func (a *Article) ClearRelLabels() {
	m2m := orm.NewOrm().QueryM2M(a, "Labels")
	m2m.Clear()
}

func (a *Article) UpdateLabels(labels ...Label) {
	m2m := orm.NewOrm().QueryM2M(a, "Labels")
	m2m.Add(labels)
}

func (a *Article) GetLabels() (labels []*Label, err error) {
	labels = make([]*Label, 0)
	id := a.Id
	num, err := orm.NewOrm().QueryTable("Label").Filter("Articles__Aticle__Id", id).Distinct().All(&labels)

	beego.Debug(num)
	return labels, err
}

//带颜色的标题
func (a *Article) ColorTitle() string {
	if a.Color != "" {
		return fmt.Sprintf("<span style=\"color:%s\">%s</span>", a.Color, a.Title)
	} else {
		return a.Title
	}
}
