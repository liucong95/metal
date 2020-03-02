package models

import (
	"time"
	"sync"

	"github.com/astaxie/beego/orm"
)

//Role ...
type Role struct {
	ID        uint      `json:"id"`
	Description    string `json:"description"`
	Privileges      string `json:"privileges"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

//UserGroups ...
type UserGroups struct {
	RoleID     uint   `json:"roleId"`
	Description string `json:"description"`
	Checked     bool   `json:"checked"`
}

func init() {
	orm.RegisterModel(new(Role))
}

// GetRolesAndUserPermission 获取所有权限和单个用户拥有的权限
func GetRolesAndUserPermission(userID int) (allRoles []Role, userRoles []uint, returnErr error) {
	o := orm.NewOrm()
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		_, err := o.Raw("SELECT id, description FROM role ORDER BY id DESC;").QueryRows(&allRoles)
		if nil != err {
			returnErr = err
		}
	}()
	go func() {
		defer wg.Done()
		_, err := o.Raw("SELECT role_id FROM privileges WHERE user_id = ? ORDER BY id DESC;", userID).QueryRows(&userRoles)
		if nil != err {
			returnErr = err
		}
	}()
	wg.Wait()
	return
}
