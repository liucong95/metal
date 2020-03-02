package controllers

import (
	"fmt"
	"strings"

	"github.com/astaxie/beego/logs"

	"metal/models"
	"metal/util"
)

//LoginController ...
type LoginController struct {
	AdminBaseController
}

//Login 。。。
func (c *LoginController) Login() {
	c.TplName = "admin/login.html"
}

//ToLogin 登录
func (c *LoginController) ToLogin() {
	var account = c.GetString("account")
	var password = c.GetString("password")

	user,err := models.GetUserByAccount(account)
	if err != nil {
		logs.Error(err)
		c.Data["json"] = ErrorData(err)
	} else if user.Status == 0 {
		c.Data["json"] = ErrorMsg("该账号已禁用，不能登录！")
	} else if user.Password != util.GetMD5(password) {
		c.Data["json"] = ErrorMsg("密码不正确！")
	} else {
		roleList, err := models.GetGroupByUserID(user.ID)
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
					loginLog := new(models.Log)
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


 //LoginOut 登出
func (c *LoginController) LoginOut() {
	c.DelSession("loginUser")
	c.Redirect("/admin/login", 302)
}