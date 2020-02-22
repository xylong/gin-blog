package model

import "time"

type ArticleInfo struct {
	Id           int64     `db:"id"`
	CategoryId   int64     `db:"category_id"`
	Summary      string    `db:"summary"`
	Title        string    `db:"title"`
	ViewCount    uint      `db:"view_count"`
	CommentCount uint      `db:"comment_count"`
	Username     string    `db:"username"`
	CreateTime   time.Time `db:"create_time"`
}

// 文章详情
type ArticleDetail struct {
	ArticleInfo
	Category
	Content string `db:"content"`
}

// 用于文章上下页
type ArticleRecord struct {
	ArticleInfo
	Category
}
