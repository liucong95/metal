package models

import (
	"github.com/astaxie/beego/orm"
)

type Authority struct {
	BaseModel
	UserId uint `json:"user_id"`
	RoleId uint `json:"role_id"`
}

func init() {
	orm.RegisterModel(new(Authority))
}

func (auth *Authority) GetUserAuthorityList(user_id uint) ([]Authority, error) {
	o := orm.NewOrm()
	var Authoritys []Authority
	num, err := o.Raw("select * from authority where user_id = ?;", user_id).QueryRows(&Authoritys)
	if nil != err && num > 0 {
		return nil, err
	}

	return Authoritys, nil
}

//GetGroupByUserId 根据userid获取usergroup list
func (auth *Authority) GetGroupByUserId(userId uint) ([]Role, error) {
	o := orm.NewOrm()
	var userGroups []Role
	//var userGroups []orm.Params//orm.Params是一个map类型
	_, err := o.Raw(`SELECT role.authority from role
		INNER JOIN authority ON role.id = authority.role_id
		WHERE authority.user_id = ?;`, userId).QueryRows(&userGroups)
	if nil != err {
		return nil, err
	}
	return userGroups, nil
}

//Save 用户添加权限
func (auth *Authority) Save() (int64, error) {
	//	var o Ormer
	o := orm.NewOrm()
	// 每次操作都需要新建一个Ormer变量，当然也可以全局设置
	// 需要 切换数据库 和 事务处理 的话，不要使用全局保存的 Ormer 对象。
	return o.Insert(auth)
}

//Delete 用户删除权限
func (auth *Authority) Delete() (int64, error) {
	//	var o Ormer
	o := orm.NewOrm()
	// 每次操作都需要新建一个Ormer变量，当然也可以全局设置
	// 需要 切换数据库 和 事务处理 的话，不要使用全局保存的 Ormer 对象。
	return o.Delete(auth)
}