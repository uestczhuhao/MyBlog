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

func (this *IndexController) Myqq() {
	this.TplName = "myblog/qq.html"
}

func (this *IndexController) Wechat() {
	this.TplName = "myblog/wechat.html"
}

func (this *IndexController) Funny() {
	this.TplName = "myblog/funny.html"
}

func (this *IndexController) Blog() {
	this.TplName = "myblog/blog.html"
}
