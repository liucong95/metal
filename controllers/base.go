package controllers

import (
	"regexp"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/logs"

	"metal/models"
)

//BaseController .
type BaseController struct {
	beego.Controller
}

//AdminBaseController ...
type AdminBaseController struct {
	BaseController
}

//UserPermission ...
type UserPermission struct {
	User       models.User
	Privileges []string //特权
}

//HasIndexPermission 前端权限验证
var HasIndexPermission = func(ctx *context.Context) {
	// TODO 做一些验证
	var i int = 10
	print(i)
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

		//暂时不需要权限
		return

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

//Result 接口返回数据标准化
type Result struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

//ErrorMsg 返回错误信息，code是可选的自定义代码
func ErrorMsg(msg string, code ...int) Result {
	var r Result
	if len(code) > 0 {
		r.Code = code[0]
	} else {
		r.Code = 1
	}
	r.Msg = msg
	r.Data = nil
	return r
}

//ErrorData ErrorMsg和ErrorData作用一样，只不过是为了方便调用方不用手动msg.Error()，只需传error类型即可
func ErrorData(msg error, code ...int) Result {
	var r Result
	if len(code) > 0 {
		r.Code = code[0]
	} else {
		r.Code = 1
	}
	r.Msg = msg.Error()
	r.Data = nil
	return r
}

//SuccessData ...
func SuccessData(data interface{}) Result {
	var r Result
	r.Code = 0
	r.Msg = "ok"
	r.Data = data
	return r
}