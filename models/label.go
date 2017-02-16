package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type Label struct {
	Id       int
	Name     string `orm:"size(20)"`
	Count    int
	Articles []*Article `orm:"rel(m2m)"`
}

func (l *Label) Add() error {
	if _, err := orm.NewOrm().Insert(l); err != nil {
		return err
	}
	return nil
}

func (l *Label) Delete() error {
	if _, err := orm.NewOrm().Delete(l); err != nil {
		return err
	}
	return nil
}

func (l *Label) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(l, fields...); err != nil {
		return err
	}
	return nil
}

func (l *Label) Read(fields ...string) error {
	if err := orm.NewOrm().Read(l, fields...); err != nil {
		return err
	}
	return nil
}

func (l *Label) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(l)
}

func (l *Label) ClearRelAtcs() {
	m2m := orm.NewOrm().QueryM2M(l, "Articles")
	m2m.Clear()
}

func (l *Label) UpdateArticles(articles ...Article) {
	m2m := orm.NewOrm().QueryM2M(l, "Articles")
	m2m.Add(articles)
}

func (l *Label) GetArticles() (articles []*Article, err error) {
	articles = make([]*Article, 0)
	id := l.Id
	num, err := orm.NewOrm().QueryTable("Article").Filter("Labels__Label__Id", id).Distinct().All(&articles)

	beego.Debug(num)

	return articles, err
}
