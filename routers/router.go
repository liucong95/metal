package routers

import (
	"metal/controllers"

	"github.com/astaxie/beego"
)

func init() {
	//如果登录过期，则返回登录界面
	beego.InsertFilter("/",beego.BeforeRouter,controllers.HasAdminPermission,false) //过滤器
	beego.Router("/",&controllers.AdminController{}, "get:Login")

	//admin管理后台路由配置
	ns := beego.NewNamespace("/admin",

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

		beego.NSRouter("/article-route", &controllers.ArticleController{}, "get:CreateArticleRoute"),
		beego.NSRouter("/article-create", &controllers.ArticleController{}, "post:CreateArticle"),
		beego.NSRouter("/articles-route", &controllers.ArticleController{}, "get:ArticlesRoute"),
		beego.NSRouter("/articles-list", &controllers.ArticleController{}, "get:ArticlesList"),
		beego.NSRouter("/article-edit-route/:id", &controllers.ArticleController{}, "get:ArticleEditRoute"),
		beego.NSRouter("/article-edit/:id", &controllers.ArticleController{}, "put:ArticleEdit"),
		beego.NSRouter("/article-delete/:id", &controllers.ArticleController{}, "delete:ArticleDelete"),
		/*
		//也可以使用注解自动路由
		beego.NSInclude(
			&controllers.GroupController{},
			&controllers.AdminController{},
			&controllers.JobCountController{},
			&controllers.OfficialController{},
		),*/
	)
	//注册namespace
	beego.AddNamespace(ns)
}
