package db

import (
	"gin-blog/model"
	"github.com/jmoiron/sqlx"
)

func InsertCategory(category *model.Category) (id int64, err error) {
	sql := "insert into category(name,no) value(?,?)"
	result, err := DB.Exec(sql, category.Name, category.No)
	if err != nil {
		return
	}
	id, err = result.LastInsertId()
	return
}

func GetCategory(id int64) (category *model.Category, err error) {
	category = &model.Category{}
	sql := "select id,name,no from category where id=?"
	err = DB.Get(category, sql, id)
	return
}

func GetCategories(ids []int64) (categories []*model.Category, err error) {
	sql, args, err := sqlx.In("select id,name,no from category where id in (?)", ids)
	if err != nil {
		return
	}
	err = DB.Select(&categories, sql, args...)
	if err != nil {
		return
	}
	return
}

func GetAllCategories() (categories []*model.Category, err error) {
	sql := "select id,name,no from category order by no"
	err = DB.Select(&categories, sql)
	return
}
