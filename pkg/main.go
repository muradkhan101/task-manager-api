package main

import (
	"fmt"

	"github.com/jmoiron/sqlx"

	"github.com/gin-gonic/gin"
)

type Login struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// DB is a connection to the AWS MySql db
var DB *sqlx.DB

func main() {
	DB = Connect()
	r := gin.Default()
	r.GET("/gj", func(c *gin.Context) {
		fmt.Println("At gj!")
		c.String(200, "<p>Good job!</p>")
	})
	test := r.Group("/test")
	{
		test.GET("/", func(c *gin.Context) {
			fmt.Println("At test root")
			c.String(200, "Wowowowow")
		})
		test.GET("/second", func(c *gin.Context) {
			fmt.Println("At test/second")
			c.JSON(200, gin.H{"parm": "chicken", "warm": "sandwich"})
		})
	}
	test2 := r.Group("/rootSecond")
	{
		test2.GET("route", func(c *gin.Context) {
			fmt.Println("rootSecond route")
			c.JSON(200, gin.H{"one-two": "buckle my show", "3-4": "shut the door"})
		})
		test2.GET("/", func(c *gin.Context) {
			fmt.Println("test2 root")
			c.String(200, "It works!")
		})
		test2.POST("/post", func(c *gin.Context) {
			var login Login
			if err := c.ShouldBindJSON(&login); err == nil {
				fmt.Println("You logged in!")
			} else {
				fmt.Println("ERROR! RED ALERT! RED ALERT!")
			}
		})
	}
	r.Run(":3000")
}

// Examples for sql

// const INSERT_TICKER = "INSERT INTO ticker_data (type, time, product_id, trade_id, sequence, price, side, last_size, best_bid, best_ask) VALUES (:type, :time, :product_id, :trade_id, :sequence, :price, :side, :last_size, :best_bid, :best_ask)"
// _, err := db.NamedExec(INSERT_TICKER, ticker)
// db, err := sqlx.Open("mysql", "username:password@tcp(url.cynquq8xdx58.us-west-2.rds.amazonaws.com)/trading_data")
// _ "github.com/go-sql-driver/mysql"
// 	"github.com/jmoiron/sqlx"
