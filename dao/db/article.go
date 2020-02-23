package db

import "gin-blog/model"

func InsertArticle(detail *model.ArticleDetail) (id int64, err error) {
	if detail == nil {
		return
	}
	sql := "insert into article(title,content,summary,username,category_id,view_count,comment_count) values(?,?,?,?,?,?,?)"
	result, err := DB.Exec(sql, detail.Title, detail.Content, detail.Summary, detail.Username, detail.CategoryId, detail.ViewCount, detail.CommentCount)
	if err != nil {
		return
	}
	id, err = result.LastInsertId()
	return
}

// GetArticleList 文章列表
func GetArticleList(pageNum, pageSize int) (articles []*model.ArticleInfo, err error) {
	if pageNum <= 0 || pageSize <= 0 {
		return
	}
	sql := "select id,category_id,title,summary,username,view_count,comment_count,create_time from article where status=1 order by create_time desc limit ?,?"
	err = DB.Select(&articles, sql, pageNum-1, pageSize)
	if err != nil {
		return
	}
	return
}

func GetArticleById(id int64) (article *model.ArticleDetail, err error) {
	if id <= 0 {
		return
	}
	sql := "select id,category_id,title,content,summary,username,view_count,comment_count,create_time from category where id=? and status=1"
	err = DB.Get(&article, sql, id)
	return
}

func GetArticleByCategory(category_id, pageNum, pageSize int) (categories []*model.ArticleInfo, err error) {
	if category_id <= 0 || pageNum <= 0 || pageSize <= 0 {
		return
	}
	sql := "select id,category_id,title,summary,username,view_count,comment_count,create_time from article where category_id=? and status=1 order by create_time desc limit ?,?"
	err = DB.Select(&categories, sql, category_id, pageNum-1, pageSize)
	return
}
