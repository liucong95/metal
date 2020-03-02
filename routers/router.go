package routers

import (
	"metal/controllers"

	"github.com/astaxie/beego"
)

func init() {
	//如果登录过期，则返回登录界面
	beego.InsertFilter("/",beego.BeforeRouter,controllers.HasAdminPermission,false) //过滤器
	beego.Router("/",&controllers.LoginController{}, "get:Login")

	//admin管理后台路由配置
	ns := beego.NewNamespace("/admin",

		//登录界面
		beego.NSRouter("/login", &controllers.LoginController{}, "get:Login"),
		//登录
		beego.NSRouter("/to-login", &controllers.LoginController{}, "post:ToLogin"),
		//登出
		beego.NSRouter("/login-out", &controllers.LoginController{}, "get:LoginOut"),

		//主界面
		beego.NSRouter("/", &controllers.AdminController{}, "get:Welcome"),
		//主界面
		beego.NSRouter("/welcome", &controllers.AdminController{}, "get:Welcome"),
		//用户列表界面
		beego.NSRouter("/user-list", &controllers.AdminController{}, "get:UserListRoute"),
		//拉取用户列表
		beego.NSRouter("/users", &controllers.AdminController{}, "get:UserList"),
		//添加用户
		beego.NSRouter("/user-add", &controllers.AdminController{}, "get:UserAddRoute"),
		beego.NSRouter("/user-add", &controllers.AdminController{}, "post:AddUser"),
		//获取用户信息 
		beego.NSRouter("/user/:id", &controllers.AdminController{}, "get:UserGet"),
		//修改用户信息
		beego.NSRouter("/user/modify", &controllers.AdminController{}, "put:UserModify"),
		//禁用用户
		beego.NSRouter("/user/forbidden", &controllers.AdminController{}, "delete,post:ForbiddenUser"),
		//获取用户权限
		beego.NSRouter("/user-roles/:id", &controllers.AuthorityController{}, "get:GetUserRoles"),
		//修改用户权限
		beego.NSRouter("/user-roles/modify", &controllers.AuthorityController{}, "post:AddUserRole"),
		
		//帖子界面
		beego.NSRouter("/article-route", &controllers.ArticleController{}, "get:CreateArticleRoute"),
		//创建帖子
		beego.NSRouter("/article-create", &controllers.ArticleController{}, "post:CreateArticle"),
		//帖子列表界面
		beego.NSRouter("/articles-route", &controllers.ArticleController{}, "get:ArticlesRoute"),
		//拉取帖子列表
		beego.NSRouter("/articles-list", &controllers.ArticleController{}, "get:ArticlesList"),
		//编辑帖子界面
		beego.NSRouter("/article-edit-route/:id", &controllers.ArticleController{}, "get:ArticleEditRoute"),
		//修改帖子
		beego.NSRouter("/article-edit/:id", &controllers.ArticleController{}, "put:ArticleEdit"),
		//删除帖子
		beego.NSRouter("/article-delete/:id", &controllers.ArticleController{}, "delete:ArticleDelete"),
	)
	//注册namespace
	beego.AddNamespace(ns)
}
