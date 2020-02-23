package controller

import (
	"gin-blog/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Index(c *gin.Context) {
	articles, err := service.GetArticles(1, 15)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
		"data":    len(articles),
	})
}
