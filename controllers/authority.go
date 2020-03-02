package controllers

import (
	"strconv"
	"time"
	"errors"
	"encoding/json"
	
	"github.com/astaxie/beego/logs"

	"metal/models"
)
// AuthorityController 用户权限管理
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
		UserID uint `json:"UserId"`
		Roles  []uint `json:"Roles"`
	}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &args); err!=nil {
		panic(err)
	}

	if args.UserID == 0 {
		panic(errors.New("userId不能为空"))
	}

	roleIds := args.Roles
	if len(roleIds) == 0 {
		panic(errors.New("roleId不能为空"))
	}

	oldRoles,err := models.GetUserAuthorityList(args.UserID)
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
			var userGroup = new(models.Authority)
			userGroup.UserId = args.UserID
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

	allRoles, userRoles, err := models.GetRolesAndUserPermission(uid)
	logs.Debug("allRoles:", allRoles)
	logs.Debug("userRoles:", userRoles)
	if nil != err {
		c.Data["json"] = ErrorData(err)
	}
	userPermissions := make([]models.UserGroups, 0, 20)
	for index, item := range allRoles {
		userPremis := new(models.UserGroups)
		userPremis.RoleID = uint(item.ID)
		userPremis.Description = item.Description
		for _, rid := range userRoles {
			if item.ID == rid {
				userPremis.Checked = true
				break
			}
		}
		userPermissions = append(userPermissions[:index], *userPremis)
	}
	c.Data["json"] = SuccessData(userPermissions)
	c.ServeJSON()
}
