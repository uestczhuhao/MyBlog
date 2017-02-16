package main

import (
	"MyBlog/models"
	_ "MyBlog/routers"
	"time"

	"github.com/astaxie/beego"
)

func init() {
	models.Init()
}

func main() {
	// var label models.Label
	// var article models.Article
	// label.Query().Filter("id", 1).One(&label)
	// article.Query().Filter("Id", 1).One(&article)
	// label.UpdateArticles(article)
	// beego.Debug(article)
	beego.Debug(time.Now())
	beego.Run()
}
