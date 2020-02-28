package controllers

//包名并非必须和文件夹名相同，但是按照惯例最后一个路径名和包名一致
import (
	"encoding/json"
	"errors"
	"github.com/astaxie/beego/logs"
	. "metal/models"
	"strconv"
	"time"
)
// GroupController 用户权限管理
type AuthorityController struct {
	AdminBaseController
}

// AddUserRole 用户添加权限
func (c *AuthorityController) AddUserRole() {
	defer func() {
		if err := recover(); err != nil {
			c.Data["json"] = ErrorData(err.(error))
		}
		c.ServeJSON()
	}()
	var args struct {
		UserId uint
		Roles  []uint
	}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &args); err!=nil {
		panic(err)
	}

	UserID := args.UserId
	if UserID == 0 {
		panic(errors.New("userId不能为空"))
	}
	roleIds := args.Roles
	if len(roleIds) == 0 {
		panic(errors.New("roleId不能为空"))
	}

	var authority Authority
	oldRoles,err := authority.GetUserAuthorityList(UserID)
	if err != nil{
		panic(err)
	}

	//添加新的
	for _,roleID := range roleIds {
		isExist := false
		for _,role := range oldRoles{
			if roleID == role.RoleId{
				isExist = true
				break
			}
		}

		//添加新的
		if !isExist{
			var userGroup = new(Authority)
			userGroup.UserId = UserID
			userGroup.RoleId = roleID
			userGroup.CreatedAt = time.Now()
			userGroup.UpdatedAt = time.Now()
			_, err := userGroup.Save()
			if nil != err {
				logs.Error(err)
				panic(err)
			}
		}
	}

	//删除旧的
	for _,role := range oldRoles {
		isExist := false
		for _,roleID := range roleIds{
			if role.RoleId == roleID{
				isExist = true
				break
			}
		}

		//删除旧的
		if !isExist{
			_, err := role.Delete()
			if nil != err {
				logs.Error(err)
				panic(err)
			}
		}
	}
	c.Data["json"] = SuccessData(nil)
}

// GetUserRoles 获取用户权限
func (c *AuthorityController) GetUserRoles() {
	uid, err := strconv.Atoi(c.Ctx.Input.Param(":id"))
	if err != nil{
		logs.Error("get uid,err:%s",err)
	}

	role := new(Role)
	allRoles, userRoles, err := role.GetRolesAndUserPermission(uid)
	logs.Debug("allRoles:", allRoles)
	logs.Debug("userRoles:", userRoles)
	if nil != err {
		c.Data["json"] = ErrorData(err)
	}
	userPermissions := make([]UserGroups, 0, 20)
	for index, item := range allRoles {
		userPremis := new(UserGroups)
		userPremis.Role_id = uint(item.Id)
		userPremis.Description = item.Description
		for _, rid := range userRoles {
			if item.Id == rid {
				userPremis.Checked = true
				break
			}
		}
		userPermissions = append(userPermissions[:index], *userPremis)
	}
	c.Data["json"] = SuccessData(userPermissions)
	c.ServeJSON()
}
