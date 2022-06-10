package main

import (
	"database/sql"
	"fmt"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
)

const (
	DB_DSN = "host=52.73.10.26 port=5432 user=crud_user password=DendronLover123! dbname=credentials sslmode=disable"
)

func get_response(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "Success bb >:)",
	})
}

func post_data(ctx *gin.Context) {
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

type User struct {
	id    int
	first string
	last  string
}

//temp db test function redo the db server using that guide
func db_test(ctx *gin.Context) {
	db, err := sql.Open("postgres", DB_DSN)
	if err != nil {
		print("Failed to open a DB connection: ", err)
	}
	err = db.Ping()
	if err != nil {
		print("\nNo open: ", err.Error())
	}
	defer db.Close()
	var first string
	var last string
	result, excerr := db.Exec("INSERT INTO users (first, last) VALUES ($1, $2)", "new", "insertion")
	if excerr != nil {
		print(excerr.Error())
	}
	print(result.LastInsertId())
	rows := db.QueryRow("SELECT first, last FROM users WHERE id =$1", 2)
	switch scanErr := rows.Scan(&first, &last); scanErr {
	case sql.ErrNoRows:
		print("\n No Rows!\n")
	case nil:
		print("successfully copied\n")
	default:
	}
	ctx.JSON(200, gin.H{"result": rows, "message": "good", "first": first, "last": last})
}

func main() {
	router := gin.Default()

	router.Use(CORSMiddleware())
	router.GET("/get_response", get_response)
	router.GET("/db_test", db_test)
	router.POST("/post_data", post_data)
	router.Run() // listen and serve on 0.0.0.0:8080
}
