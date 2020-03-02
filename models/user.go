package models

import (
	"time"
	"strconv"
	"sync"

	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/logs"
)

//SexMap 性别
var SexMap = map[int]string{0: "女", 1: "男"}

//User 用户
type User struct {
	ID        uint      `json:"id"`
	Account		string `jsong:"account"`
	Password    string `json:"password"`
	UserName    string `json:"user_name"`
	Mobile      string `json:"mobile"`
	Email       string `json:"email"`
	Description string `json:"description"`
	Status      int    `json:"status"` // 0不可用，1可用
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

//UserSt ...
type UserSt struct {
	User
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func init() {
	// 需要在init中注册定义的model
	orm.RegisterModel(new(User))
	// 如果使用 orm.QuerySeter 进行高级查询的话，这个是必须的。
	// 反之，如果只使用 Raw 查询和 map struct，是无需这一步的。
}

//Save 添加用户
func (user *User) Save() (int64, error) {
	o := orm.NewOrm()
	return o.Insert(user)
}

//Update 通过id修改用户
func (user *User) Update() (int64, error) {
	o := orm.NewOrm()
	id, err := o.Update(user, "account", "username", "email", "mobile", "description", "updated_at") // 要修改的对象和需要修改的字段
	return id, err
}
//UpdateWithPwd 通过id修改用户
func (user *User) UpdateWithPwd() (int64, error) {
	o := orm.NewOrm()
	id, err := o.Update(user, "account", "username", "password", "email", "mobile", "description", "updated_at") // 要修改的对象和需要修改的字段
	return id, err
}

//UpdateStatus 修改状态
func (user *User) UpdateStatus() (int64, error) {
	o := orm.NewOrm()
	id, err := o.Update(user, "status") // 要修改的对象和需要修改的字段
	if err != nil {
		return id, err
	}

	return id, nil
}

//GetUserByID 通过id查找用户
func GetUserByID(userID int) (*User, error) {
	user := &User{ID:uint(userID)}
	o := orm.NewOrm()
	err := o.Read(user, "id")
	return user, err
}

//GetUserByName 通过用户名查找用户
func GetUserByName(userName string) (*User,error) {
	user := &User{UserName:userName}
	o := orm.NewOrm()
	err := o.Read(user, "username")
	if err != nil {
		return nil, err
	}
	return user, nil
}

//GetUserByAccount 通过手机号查找用户
func GetUserByAccount(account string) (*User,error) {
	user := &User{Account:account}
	o := orm.NewOrm()
	err := o.Read(user, "account")
	if err != nil {
		return nil, err
	}
	return user,nil
}

//GetUserAll 获取用户列表
func GetUserAll() ([]User, error) {
	o := orm.NewOrm()
	var users []User
	num, err := o.Raw("SELECT * FROM user").QueryRows(&users)
	logs.Info("查询到", num, "条数据")
	return users, err

}

//GetUserAllByCondition 获取用户列表
func GetUserAllByCondition(cond map[string]string, start, perPage int) (users []User, total int64, newError error) {
	o := orm.NewOrm()
	var condition = " WHERE 1 "
	if cond["account"] != "" {
		condition += "and (account like '" + cond["account"] + "%' or user_name like '" + cond["username"] + "%' )"
	}
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		var sql = "SELECT * FROM user "
		sql += condition
		sql += " LIMIT ?, ?"
		_, err := o.Raw(sql, strconv.Itoa(start), strconv.Itoa(perPage)).QueryRows(&users)
		if err != nil {
			newError = err
		}
	}()
	go func() {
		defer wg.Done()
		var sql = "SELECT COUNT(0) FROM user "
		sql += condition
		err := o.Raw(sql).QueryRow(&total)
		if err != nil {
			newError = err
		}
		logs.Info("mysql row affected nums: ", total)
	}()
	wg.Wait()
	return users, total, newError
}
