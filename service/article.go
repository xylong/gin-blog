package service

import (
	"gin-blog/dao/db"
	"gin-blog/model"
)

func GetArticles(pageNum, pageSize int) (articles []*model.ArticleRecord, err error) {
	// 获取文章
	list, err := db.GetArticleList(pageNum, pageSize)
	if err != nil {
		return
	}
	if len(list) <= 0 {

		return
	}
	// 获取文章对应分类
	categoryIds := getCategoryIds(list)
	categories, err := db.GetCategories(categoryIds)
	if err != nil {
		return
	}
	// 合并
	for _, article := range list {
		articleRecord := &model.ArticleRecord{
			ArticleInfo: *article,
		}
		for _, category := range categories {
			if article.CategoryId == category.Id {
				articleRecord.Category = *category
				break
			}
		}
		articles = append(articles, articleRecord)
	}
	return
}

// 获取分类ID
func getCategoryIds(articles []*model.ArticleInfo) (ids []int64) {
	for _, article := range articles {
		cid := article.CategoryId
		// 去重
		for _, id := range ids {
			if cid != id {
				ids = append(ids, cid)
			}
		}
	}
	return
}

func GetArticlesByCategory(categoryId, pageNum, pageSize int) (articles []*model.ArticleRecord, err error) {
	list, err := db.GetArticleByCategory(categoryId, pageNum, pageSize)
	if err != nil || len(list) <= 0 {
		return
	}
	// 获取文章对应分类
	categoryIds := getCategoryIds(list)
	categories, err := db.GetCategories(categoryIds)
	if err != nil {
		return
	}
	// 合并
	for _, article := range list {
		articleRecord := &model.ArticleRecord{
			ArticleInfo: *article,
		}
		for _, category := range categories {
			if article.CategoryId == category.Id {
				articleRecord.Category = *category
				break
			}
		}
		articles = append(articles, articleRecord)
	}
	return
}
