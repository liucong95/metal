package models

import (
	"sync"
	"time"

	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/logs"
)

//Article 帖子
type Article struct {
	Id        uint      `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Status  uint8  `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

//ArticlePortal ...
type ArticlePortal struct {
	Article
	Img string
	UpdatedAt string
}

func init() {
	orm.RegisterModel(new(Article))
}

//Save ...
func (model *Article) Save() (int64, error) {
	o := orm.NewOrm()
	model.Status = 1
	model.CreatedAt = time.Now()
	model.UpdatedAt = time.Now()
	return o.Insert(model)
}

//GetArticlesByCondition ...
func GetArticlesByCondition(param map[string]string, pageIndex, pageSize int) (articles []Article, total int64, returnError error) {
	o := orm.NewOrm()
	var condition = ""
	if param["status"] != "" {
		condition += " AND status IN (" + param["status"] + ")"
	}
	if param["title"] != "" {
		condition += " AND title LIKE '" + param["title"] + "%'"
	}
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		var sql = "SELECT * FROM article WHERE 1=1"
		sql += condition
		sql += " ORDER BY id DESC"
		sql += " LIMIT ?, ?;"
		_, err := o.Raw(sql, pageIndex, pageSize).QueryRows(&articles)
		if err != nil {
			returnError = err
		}
	}()
	go func() {
		defer wg.Done()
		var sql = "SELECT COUNT(0) FROM article WHERE status = 1"
		sql += condition
		err := o.Raw(sql).QueryRow(&total)
		if err != nil {
			returnError = err
		}
		logs.Info("mysql row affected nums: ", total)
	}()
	wg.Wait()
	return articles, total, returnError
}

//GetByID ...
func (model *Article) GetByID() error {
	o := orm.NewOrm()
	err := o.Read(model, "id")
	return err
}

//Update ...
func (model *Article) Update() (int64, error) {
	o := orm.NewOrm()
	id, err := o.Update(model, "title", "content", "updated_at")
	return id, err
}

//Delete ...
func (model *Article) Delete() (int64, error) {
	o := orm.NewOrm()
	id, err := o.Delete(model)
	return id, err
}

//GetCategory ...
func (model *Article) GetCategory() ([]ArticlePortal, error) {
	o := orm.NewOrm()
	titles := make([]ArticlePortal, 1)
	_, err := o.Raw("SELECT id, title, DATE_FORMAT(updated_at,'%Y-%m-%d %H:%i') as updated_at FROM article WHERE status = 1 ORDER BY id DESC;").QueryRows(&titles)
	return titles, err
}
