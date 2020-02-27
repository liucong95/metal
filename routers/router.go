package routers

import (
	"metal/controllers"

	"github.com/astaxie/beego"
)

func init() {

	//如果登录过期，则返回登录界面
	beego.InsertFilter("/",beego.BeforeRouter,controllers.HasAdminPermission,false) //过滤器
	beego.Router("/",&controllers.AdminController{}, "get:Login")
	//namespace中的路由推荐用NS开头的方法
	//admin管理后台路由配置
	ns := beego.NewNamespace("/admin",

		//根据实际经验，推荐使用直接注册的方式来管理路由，而不是注解路由。
		//原因：可以直接明了的看到所有路由，不再使用可以直接注释掉，方便搜索。而注解路由过于分散不易管理，且不能减少代码量
		beego.NSRouter("/", &controllers.AdminController{}, "get:Welcome"),
		beego.NSRouter("/login", &controllers.AdminController{}, "get:Login"),
		beego.NSRouter("/to-login", &controllers.AdminController{}, "post:ToLogin"),
		beego.NSRouter("/login-out", &controllers.AdminController{}, "get:LoginOut"),
		beego.NSRouter("/welcome", &controllers.AdminController{}, "get:Welcome"),
		beego.NSRouter("/user-list", &controllers.AdminController{}, "get:UserListRoute"),
		beego.NSRouter("/user-add", &controllers.AdminController{}, "get:UserAddRoute"),
		beego.NSRouter("/user", &controllers.AdminController{}),
		beego.NSRouter("/user/delete", &controllers.AdminController{}, "delete,post:DeleteUser"),
		beego.NSRouter("/user/:id", &controllers.AdminController{}, "get:UserGet"),
		beego.NSRouter("/users", &controllers.AdminController{}, "get:UserList"),
		beego.NSRouter("/user-group", &controllers.UserGroupController{}),
		beego.NSRouter("/icons", &controllers.PictureController{}, "get:IconList"),
		beego.NSRouter("/picture", &controllers.PictureController{}, "get:Picture"),
		beego.NSRouter("/pictures", &controllers.PictureController{}, "post:AddPicture"),
		beego.NSRouter("/picture-list-route", &controllers.PictureController{}, "get:ListPictureRoute"),
		beego.NSRouter("/picture-list", &controllers.PictureController{}, "get:ListPicture"),
		beego.NSRouter("/picture-delete", &controllers.PictureController{}, "delete,post:DeletePicture"),

		//也可以使用注解自动路由
		beego.NSInclude(
			&controllers.GroupController{},
			&controllers.AdminController{},
			&controllers.JobCountController{},
			&controllers.WebSiteController{},
		),
	)
	//注册namespace
	beego.AddNamespace(ns)
}
