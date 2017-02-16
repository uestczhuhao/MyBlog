package models

import (
	"MyBlog/util"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

const (
	_DB_DRIVER = "mysql"
)

var _DB_CONNECT_STR string = beego.AppConfig.String("mysqluser") + ":" +
	beego.AppConfig.String("mysqlpass") + "@/" +
	beego.AppConfig.String("mysqldb") + "?charset=utf8" + "&loc=Local"

func Init() {
	//set defalut database
	orm.RegisterDataBase("default", _DB_DRIVER, _DB_CONNECT_STR)

	// register model
	orm.RegisterModel(new(User), new(Article), new(Label))

	// create table
	orm.RunSyncdb("default", false, true)

	userInit := &User{Id: 1}
	if err := userInit.Read(); err != nil {
		userInit.Id = 1
		userInit.UserName = "zhu"
		userInit.Password = util.Md5([]byte("123456"))
		userInit.Email = "uestczhuhao@163.com"
		userInit.Active = 1
		err = userInit.Add()
		if err != nil {
			beego.Debug(err)
		}
	}
}
