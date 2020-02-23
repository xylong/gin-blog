package service

import (
	"gin-blog/dao/db"
	"gin-blog/model"
)

func GetAllCategoryList(categories []*model.Category, err error) {
	categories, err = db.GetAllCategories()
	if err != nil {
		return
	}
	return
}
