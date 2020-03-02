package models

import (
	"sync"
	"time"

	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/logs"
)

//Official 官方账号
type Official struct {
	ID        uint      `json:"id"`
	Name     string `json:"name"`
	Category string `json:"category"`
	Status   int  `json:"status"`
	Url      string `json:"url"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func init() {
	orm.RegisterModel(new(Official))
}

func (model *Official) Save() (int64, error) {
	o := orm.NewOrm()
	model.Status = 1
	model.CreatedAt = time.Now()
	model.UpdatedAt = time.Now()
	return o.Insert(model)
}

func (model *Official) GetListByCondition(param map[string]string, pageIndex, pageSize int) (list []Official, total int64, returnError error) {
	o := orm.NewOrm()
	var condition = ""
	if param["status"] != "" {
		condition += " AND status IN (" + param["status"] + ")"
	}
	if param["title"] != "" {
		condition += " AND name LIKE '" + param["title"] + "%'"
	}
	list = []Official{} //初始化一个空的
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		var sql = "SELECT * FROM official WHERE 1=1"
		sql += condition
		sql += " ORDER BY id DESC"
		sql += " LIMIT ?, ?;"
		_, err := o.Raw(sql, pageIndex, pageSize).QueryRows(&list)
		if err != nil {
			returnError = err
		}
	}()
	go func() {
		defer wg.Done()
		var sql = "SELECT COUNT(0) FROM official WHERE status = 1"
		sql += condition
		err := o.Raw(sql).QueryRow(&total)
		if err != nil {
			returnError = err
		}
		logs.Info("mysql row affected nums: ", total)
	}()
	wg.Wait()
	return list, total, returnError
}

func (model *Official) GetById() error {
	o := orm.NewOrm()
	err := o.Read(model, "id")
	return err
}

func (model *Official) Update() (int64, error) {
	o := orm.NewOrm()
	id, err := o.Update(model, "title", "content", "updated_at")
	return id, err
}
func (model *Official) Delete() (int64, error) {
	o := orm.NewOrm()
	id, err := o.Delete(model)
	return id, err
}

func (model *Official) GetCategory() ([]Article, error) {
	o := orm.NewOrm()
	titles := make([]Article, 1)
	_, err := o.Raw("SELECT id, title FROM official WHERE status = 1 ORDER BY id DESC;").QueryRows(&titles)
	return titles, err
}
