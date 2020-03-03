package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

//Privilege 权限
type Privilege struct {
	Id			uint      `json:"id"`
	UserId 		uint `json:"user_id"`
	RoleId 		uint `json:"role_id"`
	CreatedAt 	time.Time `json:"created_at"`
	UpdatedAt 	time.Time `json:"updated_at"`
}

func init() {
	orm.RegisterModel(new(Privilege))
}

//Save 用户添加权限
func (auth *Privilege) Save() (int64, error) {
	o := orm.NewOrm()
	return o.Insert(auth)
}

//Delete 用户删除权限
func (auth *Privilege) Delete() (int64, error) {
	o := orm.NewOrm()
	return o.Delete(auth)
}

//GetUserPrivilegeList ...
func GetUserPrivilegeList(userID uint) ([]Privilege, error) {
	o := orm.NewOrm()
	var Privileges []Privilege
	num, err := o.Raw("select * from Privilege where user_id = ?;", userID).QueryRows(&Privileges)
	if nil != err && num > 0 {
		return nil, err
	}

	return Privileges, nil
}

//GetGroupByUserID 根据userid获取usergroup list
func GetGroupByUserID(userID uint) ([]Role, error) {
	o := orm.NewOrm()
	var userGroups []Role
	//var userGroups []orm.Params//orm.Params是一个map类型
	_, err := o.Raw(`SELECT role.privileges from role
		INNER JOIN privilege ON role.id = privilege.role_id
		WHERE privilege.user_id = ?;`, userID).QueryRows(&userGroups)
	if nil != err {
		return nil, err
	}
	return userGroups, nil
}
