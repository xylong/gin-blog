package db

import (
	"gin-blog/model"
	"testing"
)

func init() {
	dns := "root:root@tcp(localhost:3306)/blogger?parseTime=true"
	err := Init(dns)
	if err != nil {
		panic(err)
	}
}

func TestGetCategory(t *testing.T) {
	category, err := GetCategory(1)
	if err != nil {
		panic(err)
	}
	t.Logf("category:%#v", category)
}

func TestGetCategories(t *testing.T) {
	ids := []int64{1, 2, 3}
	categories, err := GetCategories(ids)
	if err != nil {
		panic(err)
	}
	for _, category := range categories {
		t.Logf("id:%d category:%#v\n", category.Id, category)
	}
}

func TestGetAllCategories(t *testing.T) {
	categories, err := GetAllCategories()
	if err != nil {
		panic(err)
	}
	for _, category := range categories {
		t.Logf("id:%d category:%#v\n", category.Id, category)
	}
}

func TestInsertCategory(t *testing.T) {
	category := &model.Category{
		Name: "辅助",
		No:   4,
	}
	id, err := InsertCategory(category)
	if err != nil {
		panic(err)
	}
	t.Logf("id=>%d\n", id)
}
