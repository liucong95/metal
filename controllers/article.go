package controllers

import (
	"strconv"
	"time"
	"log"

	"github.com/astaxie/beego/logs"
	
	"metal/models"
	"metal/service/article"
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
	article := new(models.Article)
	article.Title = title
	article.Content = content
	article.Status = 1
	article.CreatedAt = time.Now()
	article.UpdatedAt = time.Now()
	_, err := article.Save()
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
	reqType,_ := c.GetInt32("type")
	topicID,_ := c.GetUint64("topicID")
	roleID,_ := c.GetUint64("roleID")
	//个人空间
	if reqType == 0 && roleID == 0{
		logs.Error("role id cant be nil")
		c.Data["json"] = ErrorMsg("玩家ID不能为空！")
		c.ServeJSON()
		return
	}else if reqType == 3 && topicID == 0{
		//主题
		logs.Error("topic id cant be nil")
		c.Data["json"] = ErrorMsg("主题ID不能为空！")
		c.ServeJSON()
		return
	}
	start, _ := c.GetUint32("start")
	perPage, _ := c.GetUint32("perPage")

	list, total, err := article.GetMomentList(reqType,10086,roleID, topicID, start, start + perPage)
	if nil != err {
		logs.Error(err)
		c.Data["json"] = ErrorData(err)
	} else {
		data := map[string]interface{}{
			"result": list,
			"total":  total,
		}
		c.Data["json"] = SuccessData(data)
	}
	c.ServeJSON()
}

//ArticleEditRoute 编辑文章路由 @router /article-edit-route/:id [get]
func (c *ArticleController) ArticleEditRoute() {
	article := new(models.Article)
	artID, _ := strconv.Atoi(c.Ctx.Input.Param(":id"))
	article.Id = uint(artID)
	article.GetByID()
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
	article := new(models.Article)
	artID, _ := strconv.Atoi(c.Ctx.Input.Param(":id"))
	title := c.GetString("title")
	content := c.GetString("content")
	article.Id = uint(artID)
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
	article := new(models.Article)
	artID, _ := strconv.Atoi(c.Ctx.Input.Param(":id"))
	article.Id = uint(artID)
	article.Delete()
	c.Data["json"] = SuccessData(nil)
	c.ServeJSON()
}
