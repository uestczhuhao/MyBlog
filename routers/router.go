package routers

import (
	"MyBlog/controllers/admin"
	"MyBlog/controllers/myblog"

	"github.com/astaxie/beego"
)

func init() {
	//前台路由
	beego.Router("/", &myblog.IndexController{})
	beego.Router("/myblog/qq", &myblog.IndexController{}, "*:Myqq")
	beego.Router("/myblog/wechat", &myblog.IndexController{}, "*:Wechat")
	beego.Router("/myblog/funny", &myblog.IndexController{}, "*:Funny")

	//后台路由
	beego.Router("/login", &admin.LoginController{}, "get:LoginIndex;post:LoginPost")
	beego.Router("/admin", &admin.AdminIndexController{})
	beego.Router("/admin/logout", &admin.LoginController{}, "*:Loginout")
	beego.Router("/admin/login/update", &admin.LoginController{}, "*:Loginupdate")
	beego.Router("/admin/system", &admin.AdminIndexController{}, "*:System")
	//文章相关路由
	beego.Router("/admin/label", &admin.LabelController{}, "get:Label;post:Add")
	beego.Router("/admin/article/add", &admin.ArticleController{}, "*:Add")
	beego.Router("/admin/article/edit", &admin.ArticleController{}, "*:Edit")
	beego.Router("admin/article/save", &admin.ArticleController{}, "post:Save")
	beego.Router("/admin/article/list", &admin.ArticleController{}, "*:List")
	beego.Router("/admin/article/delete", &admin.ArticleController{}, "*:Delete")
	beego.Router("/admin/article/upload", &admin.ArticleController{}, "*:Upload")

	//用户相关路由
	beego.Router("/admin/user/list", &admin.UserController{}, "*:List")
	beego.Router("/admin/user/add", &admin.UserController{}, "*:Add")
	beego.Router("/admin/user/edit", &admin.UserController{}, "*:Edit")
	beego.Router("/admin/user/delete", &admin.UserController{}, "*:Delete")

}
