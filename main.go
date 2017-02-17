package main

import (
	"MyBlog/models"
	_ "MyBlog/routers"

	"github.com/astaxie/beego"
)

func init() {
	models.Init()
}

func main() {
	beego.Run()
}
