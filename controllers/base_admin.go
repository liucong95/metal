package controllers

import (
	"regexp"

	"github.com/astaxie/beego/logs"
	"metal/controllers/permissions"
	"metal/models"

	"github.com/astaxie/beego/context"
)

//AdminBaseController ...
type AdminBaseController struct {
	BaseController
}

//UserPermission ...
type UserPermission struct {
	User       models.User
	Privileges []string //特权
}

//HasAdminPermission 后台权限验证
var HasAdminPermission = func(ctx *context.Context) {
	loginUser := ctx.Input.CruSession.Get("loginUser")
	if loginUser == nil && ctx.Input.URL() != "/admin/login" && ctx.Input.URL() != "/admin/to-login" {
		ctx.Redirect(302, "/admin/login")
	}
}

//Prepare 这个函数主要是为了用户扩展用的，这个函数会在下面定义的这些 Method 方法之前执行，用户可以重写这个函数实现类似用户验证之类
func (c *AdminBaseController) Prepare() {
	// admin-user-ctrl和user-index-ctrl都继承了base-ctrl，所以都会自动执行该方法
	// 因为前端用户界面不需要权限验证，管理后台才需要
	session := c.GetSession("loginUser")
	//session为空时不进行权限验证
	if session != nil {
		userPermission := session.(*UserPermission)
		c.Data["username"] = userPermission.User.UserName
		ctrl, runMethod := c.GetControllerAndAction() // 获取controller和method
		requestPermission := ctrl + ":" + runMethod
		logs.Info(">>run-method:", ctrl+":"+runMethod)
		privileges := userPermission.Privileges

		//不需要权限
		if _,ok := permissions.InvaildPermission[requestPermission]; ok{
			return
		}

		//需要权限认证
		hasPermission := false
		for _, priPattern := range privileges {
			if priPattern == "*"{
				hasPermission = true
				break
			}

			hasPermission,_ = regexp.MatchString(priPattern, requestPermission)
			if hasPermission {
				break
			}
		}
		if !hasPermission {
			logs.Info("权限不足")
			c.Data["json"] = ErrorMsg("权限不足")
			c.ServeJSON()
		}
	}
}
