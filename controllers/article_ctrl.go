package controllers

import (
	"github.com/astaxie/beego/logs"
	"log"
	. "metal/models" // 点操作符导入的包可以省略包名直接使用公有属性和方法
	"metal/service"
	"strconv"
	"time"
)

//ArticleController 帖子
type ArticleController struct {
	AdminBaseController
}

//CreateArticleRoute 创建文章路由
func (c *ArticleController) CreateArticleRoute() {
	c.TplName = "admin/article-create.html"
}

//CreateArticle 创建文章接口
func (c *ArticleController) CreateArticle() {
	defer func() {
		if err := recover(); err != nil {
			c.Data["json"] = ErrorMsg(err.(string))
		}
		c.ServeJSON()
	}()
	title := c.GetString("title")
	if title == "" {
		log.Panic("title不能为空")
	}
	content := c.GetString("content")
	if content == "" {
		log.Panic("content不能为空")
	}
	article := new(Article)
	article.Title = title
	article.Content = content
	article.Status = 1
	article.CreatedAt = time.Now()
	article.UpdatedAt = time.Now()
	articleService := service.NewService()
	_, err := articleService.Save(article)
	if nil != err {
		logs.Error(err)
		c.Data["json"] = ErrorData(err)
	} else {
		c.Data["json"] = SuccessData(nil)
	}
	c.ServeJSON()
}


//ArticlesRoute 文章列表路由
func (c *ArticleController) ArticlesRoute() {
	c.TplName = "admin/article-list.html"
}

//ArticlesList 文章列表接口
func (c *ArticleController) ArticlesList() {
	args := c.GetString("search") // 获取所有参数
	start, _ := c.GetInt("start")
	perPage, _ := c.GetInt("perPage")
	article := new(Article)
	param := map[string]string{
		"status": "1,0",
		"title":  args,
	}

	list, total, err := article.GetArticlesByCondition(param, start, perPage)
	if nil != err {
		logs.Error(err)
		c.Data["json"] = ErrorData(err)
	} else {
		data := map[string]any{
			"result": list,
			"total":  total,
		}
		c.Data["json"] = SuccessData(data)
	}
	c.ServeJSON()
}

//ArticleEditRoute 编辑文章路由 @router /article-edit-route/:id [get]
func (c *ArticleController) ArticleEditRoute() {
	article := new(Article)
	artId, _ := strconv.Atoi(c.Ctx.Input.Param(":id"))
	article.Id = uint(artId)
	article.GetById()
	c.Data["article"] = article
	c.TplName = "admin/article-edit.html"
}

//ArticleEdit 修改文章接口 @router /article/:id [put]
func (c *ArticleController) ArticleEdit() {
	defer func() {
		if err := recover(); err != nil {
			c.Data["json"] = ErrorData(err.(error))
			c.ServeJSON()
		}
	}()
	article := new(Article)
	artId, _ := strconv.Atoi(c.Ctx.Input.Param(":id"))
	title := c.GetString("title")
	content := c.GetString("content")
	article.Id = uint(artId)
	article.Title = title
	article.Content = content
	article.UpdatedAt = time.Now()
	_, err := article.Update()
	if err != nil {
		c.Data["json"] = ErrorData(err.(error))
		c.ServeJSON()
		return
	}
	c.Data["json"] = SuccessData(nil)
	c.ServeJSON()
}


//ArticleDelete 删除文章接口
func (c *ArticleController) ArticleDelete() {
	article := new(Article)
	artId, _ := strconv.Atoi(c.Ctx.Input.Param(":id"))
	article.Id = uint(artId)
	article.Delete()
	c.Data["json"] = SuccessData(nil)
	c.ServeJSON()
}
