package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/logs"
	. "metal/models" // 点操作符导入的包可以省略包名直接使用公有属性和方法
	"metal/util"
	"strconv"
	"strings"
	"time"
)

//AdminController ...
type AdminController struct {
	AdminBaseController
}

//Login 。。。
func (c *AdminController) Login() {
	c.TplName = "admin/login.html"
}

//ToLogin 登录
func (c *AdminController) ToLogin() {
	var account = c.GetString("account")
	var password = c.GetString("password")

	user := &User{Account: account}
	err := user.GetByAccount()
	if err != nil {
		logs.Error(err)
		c.Data["json"] = ErrorData(err)
	} else if user.Status == 0 {
		c.Data["json"] = ErrorMsg("该账号已禁用，不能登录！")
	} else if user.Password != util.GetMD5(password) {
		c.Data["json"] = ErrorMsg("密码不正确！")
	} else {
		group := new(Authority)
		roleList, err := group.GetGroupByUserId(user.Id)
		if err != nil {
			c.Data["json"] = ErrorData(err)
		}
		var privileges []string
		for _, v := range roleList {
			strArr := strings.Split(v.Authority, ",")
			privileges = append(privileges, strArr...)
		}
		userPermission := new(UserPermission)
		userPermission.User = *user
		userPermission.Privileges = privileges
		c.SetSession("loginUser", userPermission)
		// c.Ctx.Input.IP()获取到的是Nginx内网ip，需要在Nginx配置proxy_set_header Remote_addr $remote_addr;
		ip := c.Ctx.Input.Header("Remote_addr")
		// ip = "103.14.252.249"
		if ip != "" {
			//第三方接口不稳定，会阻塞整体，所以放到协程中
			go func() {
				ipBody := new(util.IPBody)
				err := util.GetIpGeography(ip, ipBody)
				if err == nil {
					loginLog := new(Log)
					// mark := fmt.Sprintf("登录IP:%s，物理地址：%s %s %s %s", ip, ipBody.Data.Country, ipBody.Data.Area, ipBody.Data.Region, ipBody.Data.City)
					mark := fmt.Sprintf("登录IP:%s，物理地址：%s", ip, ipBody.Content.Address)
					loginLog.Save(mark)
				}
			}()
		}
		c.Data["json"] = SuccessData(nil)
	}
	c.ServeJSON()
}

/**
 * 登出
 */
func (c *AdminController) LoginOut() {
	c.DelSession("loginUser")
	c.Redirect("/admin/login", 302)
}

func (c *AdminController) Welcome() {
	c.TplName = "admin/index.html"
}

func (c *AdminController) UserAddRoute() {
	c.TplName = "admin/user-add.html"
}

/**
 * 新建用户
 */
func (c *AdminController) Post() {
	args := map[string]string{}
	body := c.Ctx.Input.RequestBody
	if err := json.Unmarshal(body, &args); err != nil{
		logs.Error(err)
		c.Data["json"] = ErrorData(err)
	}else{
		var user = new(User)
		user.Account = args["account"]
		user.UserName = args["username"]
		user.Mobile = args["mobile"]
		user.Email = args["email"]
		user.Description = args["description"]
		user.Password = util.GetMD5(args["password"])
		user.CreatedAt = time.Now()
		user.UpdatedAt = time.Now()

		id, err := user.Save()
		if nil != err {
			logs.Error(err)
			c.Data["json"] = ErrorData(err)
		} else {
			c.Data["json"] = SuccessData(id)
		}
	}
	c.ServeJSON()
}

/**
 * 通过如下方式获取路由参数
 * /admin/user/:id
 * c.Ctx.Input.Param(":id")
 */
func (c *AdminController) UserGet() {
	idstr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idstr)
	user := new(User)
	user.Id = uint(id)
	userObj, err := user.GetById()
	if err != nil {
		logs.Error(err)
		c.Data["json"] = ErrorData(err)
	}
	c.Data["json"] = SuccessData(userObj)
	c.ServeJSON()
}

/**
 * 修改用户
 */
func (c *AdminController) Put() {
	userID, err := c.GetInt("userId")
	if err != nil{
		logs.Error("user id err:%s", err)
		c.Data["json"] = ErrorData(err)
	}

	var user = new(User)
	user.Id = uint(userID)
	user.Account = c.GetString("account")
	user.UserName = c.GetString("username")
	user.Email = c.GetString("email")
	user.Mobile = c.GetString("mobile")
	user.Description = c.GetString("description")
	user.UpdatedAt = time.Now()
	upID, err := user.Update()
	if nil != err {
		logs.Error(err)
		c.Data["json"] = ErrorData(err)
	} else {
		c.Data["json"] = SuccessData(upID)
	}
	c.ServeJSON()
}

/**
 * 用户列表路由
 */
func (c *AdminController) UserListRoute() {
	c.Data["Title"] = "用户列表"
	c.TplName = "admin/user-list.html"
}

/**
 * 用户列表接口
 * /admin/users
 */
func (c *AdminController) UserList() {
	args := c.GetString("search") //搜索框
	start, _ := c.GetInt("start")
	perPage, _ := c.GetInt("perPage")
	user := new(User)
	var userList = make([]UserSt, 10)
	var param = make(map[string]string)
	param["account"] = args
	param["username"] = args
	list, total, err := user.GetAllByCondition(param, start, perPage)
	if nil != err {
		logs.Error(err)
		c.Data["json"] = ErrorData(err)
	} else {
		for index, u := range list {
			userSt := new(UserSt)
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

/**
 * 删除用户
 */
func (c *AdminController) DeleteUser() {
	id, _ := c.GetInt("userId")
	var user = new(User)
	user.Id = uint(id)
	user.Status = 0
	id64, err := user.Delete()
	if nil != err {
		logs.Error(err)
		c.Data["json"] = ErrorData(err)
	} else {
		c.Data["json"] = SuccessData(id64)
	}
	c.ServeJSON()
}

/**
 * 日志列表路由
 */
// @router /logs-route
func (c *AdminController) LogsRoute() {
	c.Data["Title"] = "日志列表"
	c.TplName = "admin/logs.html"
}

/**
 * 日志列表接口
 */
// @router /logs [get]
func (c *AdminController) GetLogs() {
	start, _ := c.GetInt("start")
	perPage, _ := c.GetInt("perPage")
	var logModel = new(Log)
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