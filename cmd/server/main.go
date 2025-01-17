package main

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/task-manager-api/internal/backend"
)

// DB is a connection to the AWS MySql db
var DB *sqlx.DB

func main() {
	DB = backend.GetDb()
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.Use(jwtAuth())
	r.Use(addCors())
	r.GET("/graphql", func(c *gin.Context) {
		jwtToken, exists := c.Get("user")
		if exists {
			if jwtToken.(jwt.Token).Valid {
				query, _ := c.GetQuery("query")
				result := backend.ExecuteQuery(query)
				c.JSON(200, result)
			} else {
				c.AbortWithStatusJSON(403, gin.H{"error": "Your JWT token is invalid"})
			}
		} else {
			c.AbortWithStatusJSON(403, gin.H{"error": "Your JWT token doesn't exist"})
		}
	})
	r.OPTIONS("/graphql", optionsAdder("GET,OPTIONS"))
	user := r.Group("/user")
	{
		user.POST("/login", backend.LoginHandler)
		user.OPTIONS("/login", optionsAdder("POST,OPTIONS"))
		user.POST("/create", backend.CreateUserHandler)
		user.OPTIONS("/create", optionsAdder("POST,OPTIONS"))
	}
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
	r.Run()
}

func optionsAdder(optionTypes string) func(c *gin.Context) {
	return func(c *gin.Context) {
		c.SetAccepted(optionTypes)
		c.Status(200)
	}
}

func jwtAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if len(token) != 0 {
			// Bearer <- len is 7
			jwtToken := token[7:len(token)]
			jwtT := backend.ValidateJwt(jwtToken)
			if jwtT != nil {
				c.Set("user", *jwtT)
			}
		}
		c.Next()
	}
}

func addCors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Authorization,Content-Type")
	}
}
