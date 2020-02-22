package db

import (
	"gin-blog/model"
	"testing"
	"time"
)

func init() {
	dns := "root:root@tcp(localhost:3306)/blogger?parseTime=true"
	err := Init(dns)
	if err != nil {
		panic(err)
	}
}

func TestInsertArticle(t *testing.T) {
	article := &model.ArticleDetail{}
	article.ArticleInfo.CategoryId = 1
	article.ArticleInfo.Title = "测试"
	article.ArticleInfo.Summary = "test"
	article.Content = "this is a test"
	article.ArticleInfo.ViewCount = 1
	article.ArticleInfo.CommentCount = 0
	article.ArticleInfo.Username = "静静"
	article.ArticleInfo.CreateTime = time.Now()
	id, err := InsertArticle(article)
	if err != nil {
		panic(err)
	}
	t.Logf("%d", id)
}

func TestGetArticleList(t *testing.T) {
	rows, err := GetArticleList(1, 10)
	if err != nil {
		panic(err)
	}
	t.Logf("%d", len(rows))
}
