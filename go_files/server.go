package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	router := gin.Default()

	router.LoadHTMLGlob("../frontend/template/*")

	//Loads things from the static dir
	router.Static("../frontend/static", "../frontend/static")

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
		//hash the creds
		user, user_err := bcrypt.GenerateFromPassword([]byte(ctx.PostForm("user")), bcrypt.DefaultCost)
		pass, pass_err := bcrypt.GenerateFromPassword([]byte(ctx.PostForm("pass")), bcrypt.DefaultCost)
		//check for errors
		if user_err != nil {
			panic(user_err)
		} else if pass_err != nil {
			panic(pass_err)
		}
		fmt.Printf("first: %s \n last: %s\n", user, pass)
		ctx.JSON(200, gin.H{"message": "congrats, it works"})
	})
	router.Run() // listen and serve on 0.0.0.0:8080
}
