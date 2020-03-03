package controllers

import (
	"encoding/json"
	"strconv"
	"time"

	"github.com/astaxie/beego/logs"

	"metal/models"
	"metal/util"
)

//AdminController ...
type AdminController struct {
	AdminBaseController
}

//UserAddRoute 新建用户
func (c *AdminController) UserAddRoute() {
	c.TplName = "admin/user-add.html"
}

//UserListRoute 用户列表路由
func (c *AdminController) UserListRoute() {
	c.Data["Title"] = "用户列表"
	c.TplName = "admin/user-list.html"
}

//LogsRoute 日志列表路由
func (c *AdminController) LogsRoute() {
	c.Data["Title"] = "日志列表"
	c.TplName = "admin/logs.html"
}

//AddUser 新建用户
func (c *AdminController) AddUser() {
	args := map[string]string{}
	body := c.Ctx.Input.RequestBody
	if err := json.Unmarshal(body, &args); err != nil{
		logs.Error(err)
		c.Data["json"] = ErrorData(err)
		c.ServeJSON()
		return
	}

	account := args["account"]
	if account == "" {
		c.Data["json"] = ErrorMsg("account cant be null!")
		c.ServeJSON()
		return
	}

	passwd := args["password"]
	if passwd == "" {
		c.Data["json"] = ErrorMsg("password cant be null!")
		c.ServeJSON()
		return
	}

	name := args["username"]
	if passwd == "" {
		c.Data["json"] = ErrorMsg("username cant be null!")
		c.ServeJSON()
		return
	}

	var user = new(models.User)
	user.Account = account
	user.Password = util.GetMD5(passwd)
	user.UserName = name
	user.Mobile = args["mobile"]
	user.Email = args["email"]
	user.Description = args["description"]
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	id, err := user.Save()
	if nil != err {
		logs.Error(err)
		c.Data["json"] = ErrorData(err)
	} else {
		c.Data["json"] = SuccessData(id)
	}
	
	c.ServeJSON()
}

//UserGet ...
func (c *AdminController) UserGet() {
	id, err := strconv.Atoi(c.Ctx.Input.Param(":id"))
	if err != nil{
		logs.Error(err)
		c.Data["json"] = ErrorData(err)
	}else{
		user,err := models.GetUserByID(id)
		if err != nil {
			logs.Error(err)
			c.Data["json"] = ErrorData(err)
		}else{
			c.Data["json"] = SuccessData(user)
		}
	}
	c.ServeJSON()
}

//UserModify 修改用户
func (c *AdminController) UserModify() {
	userID, err := c.GetInt("userId")
	if err != nil{
		logs.Error("user id err:%s", err)
		c.Data["json"] = ErrorData(err)
	}

	var user = new(models.User)
	user.Id = uint(userID)
	user.Account = c.GetString("account")
	user.UserName = c.GetString("username")
	user.Email = c.GetString("email")
	user.Mobile = c.GetString("mobile")
	user.Description = c.GetString("description")
	user.UpdatedAt = time.Now()

	var upID int64 = 0
	password := c.GetString("password")
	if len(password) != 0{
		user.Password = util.GetMD5(password)
		upID, err = user.UpdateWithPwd()
	}else{
		upID, err = user.Update()
	}

	if nil != err {
		logs.Error(err)
		c.Data["json"] = ErrorData(err)
	} else {
		c.Data["json"] = SuccessData(upID)
	}
	c.ServeJSON()
}

 //UserList 用户列表接口
func (c *AdminController) UserList() {
	args := c.GetString("search") //搜索框
	start, _ := c.GetInt("start")
	perPage, _ := c.GetInt("perPage")
	var userList = make([]models.UserSt, 10)
	var param = make(map[string]string)
	param["account"] = args
	param["username"] = args
	list, total, err := models.GetUserAllByCondition(param, start, perPage)
	if nil != err {
		logs.Error(err)
		c.Data["json"] = ErrorData(err)
	} else {
		for index, u := range list {
			userSt := new(models.UserSt)
			userSt.User = u
			userSt.CreatedAt = u.CreatedAt.Format("2006-01-02 15:04:05")
			userSt.UpdatedAt = u.UpdatedAt.Format("2006-01-02 15:04:05")
			userList = append(userList[:index], *userSt)
		}
		data := map[string]interface{}{
			"result": userList,
			"total":  total,
		}
		c.Data["json"] = SuccessData(data)
	}
	c.ServeJSON()
}

//ForbiddenUser 禁用/启用
func (c *AdminController) ForbiddenUser() {
	id, _ := c.GetInt("userId")
	session := c.GetSession("loginUser")
	userPermission := session.(*UserPermission)
	if id == int(userPermission.User.Id) {
		c.Data["json"] = ErrorMsg("不能禁用自己")
		c.ServeJSON()
		return
	}
	
	user,err := models.GetUserByID(id)
	if err != nil{
		logs.Error("get user err:%s, id:%d", err, id)
		c.Data["json"] = ErrorData(err)
		c.ServeJSON()
		return
	}

	if user.Status == 0{
		user.Status = 1
	}else{
		user.Status = 0
	}
	
	_, err = user.UpdateStatus()
	if nil != err {
		logs.Error(err)
		c.Data["json"] = ErrorData(err)
	} else {
		c.Data["json"] = SuccessData(user.Status)
	}
	c.ServeJSON()
}

//GetLogs 日志列表接口
func (c *AdminController) GetLogs() {
	start, _ := c.GetInt("start")
	perPage, _ := c.GetInt("perPage")
	var logModel = new(models.Log)
	list, total, err := logModel.GetLogs(start, perPage)
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