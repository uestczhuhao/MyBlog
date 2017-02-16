package myblog

import . "MyBlog/controllers"

type IndexController struct {
	BaseController
}

func (this *IndexController) Get() {
	this.TplName = "myblog/index.html"
}

func (this *IndexController) Post() {
	this.TplName = "myblog/index.html"
}
