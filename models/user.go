package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

///用户表模型
type User struct {
	Id         int
	UserName   string    `orm:"unique;size(20)"`
	Password   string    `orm:"size(32)"`
	Email      string    `orm:"size(50)"`
	LastLogin  time.Time `orm:"auto_now_add;type(datetime)"`
	LoginCount int
	Active     int8
}

func (u *User) Add() error {
	if _, err := orm.NewOrm().Insert(u); err != nil {
		return err
	}
	return nil
}

func (u *User) Delete() error {
	if _, err := orm.NewOrm().Delete(u); err != nil {
		return err
	}
	return nil
}

func (u *User) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(u, fields...); err != nil {
		return err
	}
	return nil
}

func (u *User) Read(fields ...string) error {
	if err := orm.NewOrm().Read(u, fields...); err != nil {
		return err
	}
	return nil
}

func (u *User) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(u)
}
