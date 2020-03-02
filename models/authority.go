package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

//Authority 权限
type Authority struct {
	Id			uint      `json:"id"`
	UserId 		uint `json:"user_id"`
	RoleId 		uint `json:"role_id"`
	CreatedAt 	time.Time `json:"created_at"`
	UpdatedAt 	time.Time `json:"updated_at"`
}

func init() {
	orm.RegisterModel(new(Authority))
}

//Save 用户添加权限
func (auth *Authority) Save() (int64, error) {
	o := orm.NewOrm()
	return o.Insert(auth)
}

//Delete 用户删除权限
func (auth *Authority) Delete() (int64, error) {
	o := orm.NewOrm()
	return o.Delete(auth)
}

//GetUserAuthorityList ...
func GetUserAuthorityList(userID uint) ([]Authority, error) {
	o := orm.NewOrm()
	var Authoritys []Authority
	num, err := o.Raw("select * from authority where user_id = ?;", userID).QueryRows(&Authoritys)
	if nil != err && num > 0 {
		return nil, err
	}

	return Authoritys, nil
}

//GetGroupByUserID 根据userid获取usergroup list
func GetGroupByUserID(userID uint) ([]Role, error) {
	o := orm.NewOrm()
	var userGroups []Role
	//var userGroups []orm.Params//orm.Params是一个map类型
	_, err := o.Raw(`SELECT role.authority from role
		INNER JOIN authority ON role.id = authority.role_id
		WHERE authority.user_id = ?;`, userID).QueryRows(&userGroups)
	if nil != err {
		return nil, err
	}
	return userGroups, nil
}
