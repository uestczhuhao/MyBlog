package admin

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"

	"MyBlog/models"
	"MyBlog/util"
)

type ArticleController struct {
	AdminIndexController
}

func (this *ArticleController) List() {
	var (
		page       int
		pagesize   int = 10
		status     int
		offset     int
		list       []*models.Article
		post       models.Article
		searchtype string
		keyword    string
	)

	searchtype = this.GetString("searchtype")
	keyword = this.GetString("keyword")
	status, _ = this.GetInt("status")
	if page, _ = this.GetInt("page"); page < 1 {
		page = 1
	}
	offset = (page - 1) * pagesize

	query := post.Query().Filter("status", status)

	if keyword != "" {
		switch searchtype {
		case "title":
			query = query.Filter("title__icontains", keyword)
		case "author":
			query = query.Filter("author__icontains", keyword)
			// case "tag":
			// 	query = query.Filter("tags__icontains", keyword)
		}
	}
	count, _ := query.Count()
	if count > 0 {
		query.OrderBy("-is_top", "-post_time").Limit(pagesize, offset).All(&list)
		for _, v := range list {
			orm.NewOrm().LoadRelated(v, "Labels")
			// beego.Debug(reflect.TypeOf(v))
		}
	}

	this.Data["searchtype"] = searchtype
	this.Data["keyword"] = keyword
	this.Data["count_1"], _ = post.Query().Filter("status", 1).Count()
	this.Data["count_2"], _ = post.Query().Filter("status", 2).Count()
	this.Data["status"] = status
	this.Data["count"] = count
	this.Data["list"] = list
	this.Data["pagebar"] = util.NewPager(page, int(count), pagesize, fmt.Sprintf("/admin/article/list?status=%d&searchtype=%s&keyword=%s", status, searchtype, keyword), true).ToString()
	this.Layout = "admin/layout.html"
	this.TplName = "admin/article/list.html"
}

func (this *ArticleController) Add() {
	this.Data["posttime"] = time.Now().Format("2006-01-02 15:04:05")
	label := new(models.Label)
	num, _ := label.Query().Count()
	labels := make([]*models.Label, num)
	label.Query().All(&labels)
	this.Data["labels"] = labels

	this.Layout = "admin/layout.html"
	this.TplName = "admin/article/add.html"

}

func (this *ArticleController) Edit() {
	id, _ := this.GetInt("id")
	article := models.Article{Id: id}
	if article.Read() != nil {
		this.Abort("404")
	}

	var labels []*models.Label
	new(models.Label).Query().All(&labels)
	this.Data["labels"] = labels
	// orm.NewOrm().LoadRelated(&article, "Labels")
	// this.Data["labels"] = article.Labels
	this.Data["post"] = article
	this.Data["posttime"] = article.PostTime.Format("2006-01-02 15:04:05")
	this.Layout = "admin/layout.html"
	this.TplName = "admin/article/edit.html"

}

func (this *ArticleController) Save() {
	var (
		id      int      = 0
		title   string   = strings.TrimSpace(this.GetString("title"))
		content string   = this.GetString("content")
		labels  []string = this.GetStrings("labels")
		color   string   = strings.TrimSpace(this.GetString("color"))
		// timestr string   = strings.TrimSpace(this.GetString("posttime"))
		status  int  = 0
		istop   int8 = 0
		article models.Article
	)
	if title == "" {
		this.Showmsg("标题不能为空！")
	}

	id, _ = this.GetInt("id")
	status, _ = this.GetInt("status")

	if this.GetString("istop") == "1" {
		istop = 1
	}
	if status != 1 && status != 2 {
		status = 0
	}

	if id < 1 {
		article.UserId = this.Userid
		article.Author = this.Username
		article.PostTime = time.Now().UTC()
		article.UpdateTime = time.Now().UTC()
		article.Add()
	} else {
		article.Id = id
		if article.Read() != nil {
			goto RD
		}
	}

	article.ClearRelLabels()
	article.Status = int8(status)
	article.Title = title
	article.Color = color
	article.IsTop = istop
	article.Content = content
	article.UpdateTime = time.Now().UTC()
	article.Update("Status", "Title", "Color", "IsTop", "Content", "UpdateTime")

	if len(labels) != 0 {
		labelId := make([]int, len(labels))
		for i, tempLabel := range labels {
			labelId[i], _ = strconv.Atoi(tempLabel)
		}

		for _, tempLabel := range labelId {
			var label models.Label
			label.Query().Filter("Id", tempLabel).One(&label)
			label.Count = label.Count + 1
			label.Update("Count")
			article.UpdateLabels(label)
		}
	}
RD:
	this.Redirect("/admin/article/list", 302)
}

func (this *ArticleController) Delete() {
	id, _ := this.GetInt("id")
	article := models.Article{Id: id}
	if article.Read() == nil {
		article.ClearRelLabels()
		article.Delete()
	}
	this.Redirect("/admin/article/list", 302)
}

//上传文件
func (this *ArticleController) Upload() {
	_, header, err := this.GetFile("upfile")
	ext := strings.ToLower(header.Filename[strings.LastIndex(header.Filename, "."):])
	out := make(map[string]string)
	out["url"] = ""
	out["fileType"] = ext
	out["original"] = header.Filename
	out["state"] = "SUCCESS"
	if err != nil {
		out["state"] = err.Error()
	} else {
		savepath := "./static/upload/" + time.Now().Format("20060102")
		if err := os.MkdirAll(savepath, os.ModePerm); err != nil {
			out["state"] = err.Error()
		} else {
			filename := fmt.Sprintf("%s/%d%s", savepath, time.Now().UnixNano(), ext)
			if err := this.SaveToFile("upfile", filename); err != nil {
				out["state"] = err.Error()
			} else {
				out["url"] = filename[1:]
			}
		}
	}

	this.Data["json"] = out
	this.ServeJSON()
}
