package controllers

import (
	_ "strconv"
	"encoding/json"

	"github.com/astaxie/beego/logs"
	
	"metal/service/official"
)

//OfficialController ...
type OfficialController struct {
	AdminBaseController
}

//OfficialController 官方列表
func (c *OfficialController) OfficialListRoute() {
	c.TplName = "admin/official-list.html"
}

//AddOfficialUserRoute ...
func (c *OfficialController) AddOfficialUserRoute() {
	c.TplName = "admin/official-add.html"
}

//AddOfficialUser ...
func (c *OfficialController) AddOfficialUser() {
	args := map[string]string{}
	body := c.Ctx.Input.RequestBody
	if err := json.Unmarshal(body, &args); err != nil{
		logs.Error(err)
		c.Data["json"] = ErrorData(err)
		c.ServeJSON()
		return
	}

	name := args["username"]
	if name == ""{
		c.Data["json"] = ErrorMsg("name cant be null")
		c.ServeJSON()
		return
	}

	head := args["head"]
	signature := args["signature"]

	_, err := official.CreateOfficialAcount(name, head, signature)
	if err != nil{
		logs.Error(err)
		c.Data["json"] = ErrorData(err)
	}else{
		c.Data["json"] = SuccessData(nil)
	}
	c.ServeJSON()
}

//GetOfficialList ...
func (c *OfficialController) GetOfficialList() {
	list, total, err := official.GetOfficialAcountList()
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

//GetOfficialInfo ...
func (c *OfficialController) GetOfficialInfo() {
}

//ModifyOfficialUserInfo ...
func (c *OfficialController) ModifyOfficialUserInfo() {
	
}

//ForbiddenOfficialUser ...
func (c *OfficialController) ForbiddenOfficialUser() {
}