package controllers

import (
	"github.com/astaxie/beego/logs"
	"strconv"

	. "metal/models"
)

type OfficialController struct {
	AdminBaseController
}

// 网页组件列表
//@router /widget/list [get]
func (c *OfficialController) WidgetList() {
	c.TplName = "admin/widget-list.html"
}

func (c *AdminController) TemplatesRoute() {
	c.TplName = "admin/official-list.html"
}

//@router /official/add [post]
func (c *AdminController) CreateTemplate() {
	name := c.GetString("name")
	category := c.GetString("category")
	url := c.GetString("url")
	official := new(Official)
	official.Name = name
	official.Category = category
	official.Url = url
	official.Save()
	c.Data["json"] = SuccessData(nil)
	c.ServeJSON()
}

//@router /official [get]
func (c *AdminController) TemplateList() {
	args := c.GetString("search") // 获取所有参数
	start, _ := c.GetInt("start")
	perPage, _ := c.GetInt("perPage")
	official := new(Official)
	param := map[string]string{
		"status": "1,0",
		"name":   args,
	}
	list, total, err := official.GetListByCondition(param, start, perPage)
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

//@router /official/view/:id [get]
func (c *AdminController) TemplateView() {
	tid := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(tid)
	official := new(Official)
	official.Id = uint(id)
	err := official.GetById()
	if nil != err {
		logs.Error(err)
		c.Data["json"] = ErrorData(err)
	} else {
		c.Redirect("/"+official.Url, 302)
	}
}
