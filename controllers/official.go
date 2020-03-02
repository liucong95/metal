package controllers

import (
	"strconv"

	"github.com/astaxie/beego/logs"
	
	"metal/models"
)

//OfficialController ...
type OfficialController struct {
	AdminBaseController
}

//WidgetList 网页组件列表
func (c *OfficialController) WidgetList() {
	c.TplName = "admin/widget-list.html"
}

//TemplatesRoute ...
func (c *OfficialController) TemplatesRoute() {
	c.TplName = "admin/official-list.html"
}

//CreateTemplate ...
func (c *OfficialController) CreateTemplate() {
	name := c.GetString("name")
	category := c.GetString("category")
	url := c.GetString("url")
	official := new(models.Official)
	official.Name = name
	official.Category = category
	official.Url = url
	official.Save()
	c.Data["json"] = SuccessData(nil)
	c.ServeJSON()
}

//TemplateList ...
func (c *OfficialController) TemplateList() {
	args := c.GetString("search") // 获取所有参数
	start, _ := c.GetInt("start")
	perPage, _ := c.GetInt("perPage")
	official := new(models.Official)
	param := map[string]string{
		"status": "1,0",
		"name":   args,
	}
	list, total, err := official.GetListByCondition(param, start, perPage)
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

//TemplateView ...
func (c *OfficialController) TemplateView() {
	tid := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(tid)
	official := new(models.Official)
	official.ID = uint(id)
	err := official.GetById()
	if nil != err {
		logs.Error(err)
		c.Data["json"] = ErrorData(err)
	} else {
		c.Redirect("/"+official.Url, 302)
	}
}
