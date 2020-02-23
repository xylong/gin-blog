package main

import (
	"gin-blog/controller"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", controller.Index)
	r.Run()
}
