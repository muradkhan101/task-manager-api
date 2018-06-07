package main

import (
	"fmt"

	"github.com/jmoiron/sqlx"

	"github.com/gin-gonic/gin"
	"github.com/task-manager-api/internal/gql"
)

// DB is a connection to the AWS MySql db
var DB *sqlx.DB

func main() {
	DB = gql.Connect()
	r := gin.Default()
	r.Any("/graphql", func(c *gin.Context) {
		query, _ := c.GetQuery("query")
		fmt.Print(query)
		result := gql.ExecuteQuery(query)
		c.JSON(200, result)
	})
	// test := r.Group("/test")
	// {
	// 	test.GET("/", func(c *gin.Context) {
	// 		fmt.Println("At test root")
	// 		c.String(200, "Wowowowow")
	// 	})
	// 	test.GET("/second", func(c *gin.Context) {
	// 		fmt.Println("At test/second")
	// 		c.JSON(200, gin.H{"parm": "chicken", "warm": "sandwich"})
	// 	})
	// }
	// test2 := r.Group("/rootSecond")
	// {
	// 	test2.GET("route", func(c *gin.Context) {
	// 		fmt.Println("rootSecond route")
	// 		c.JSON(200, gin.H{"one-two": "buckle my show", "3-4": "shut the door"})
	// 	})
	// 	test2.GET("/", func(c *gin.Context) {
	// 		fmt.Println("test2 root")
	// 		c.String(200, "It works!")
	// 	})
	// }
	r.Run(":80")
}
