package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.LoadHTMLFiles("index.html")

	router.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.html", gin.H{
			"title":   "ASCIIHUB",
			"btnName": "Click this!",
		})
	})

	router.GET("/get_response", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Success bb >:)",
		})
	})

	router.POST("/post_data", func(ctx *gin.Context) {
		first := ctx.PostForm("first")
		last := ctx.PostForm("last")
		fmt.Printf("first: %s, last: %s\n", first, last)
		ctx.Redirect(http.StatusFound, "/")
	})
	router.Run() // listen and serve on 0.0.0.0:8080
}
