package main

import (
	"database/sql"
	"fmt"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

const (
	// TODO fill this in directly or through environment variable
	// Build a DSN e.g. postgres://username:password@url.com:5432/dbName
	DB_DSN = "postgres://admin:DendronLover123!@ec2-3-83-140-90.compute-1.amazonaws.com:5432/credentials"
)

func get_response(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "Success bb >:)",
	})
}

func post_data(ctx *gin.Context) {
	db, err := sql.Open("postgres", DB_DSN)
	if err != nil {
		print("Failed to open a DB connection: ", err)
	}
	defer db.Close()
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
}

//CORs set header
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func main() {
	router := gin.Default()

	router.Use(CORSMiddleware())
	router.GET("/get_response", get_response)
	router.POST("/post_data", post_data)
	router.Run() // listen and serve on 0.0.0.0:8080
}
