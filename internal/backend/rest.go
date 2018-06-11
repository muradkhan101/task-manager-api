package backend

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	. "github.com/task-manager-api/internal/types"
)

// CreateUserHandler handles creation of a new user
func CreateUserHandler(c *gin.Context) {
	var user User
	err := c.BindJSON(&user)
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
		return
	}
	query := fmt.Sprintf(GetUserByEmail, user.Email)
	var results []User
	err = GetDb().Select(&results, query)

	if len(results) > 0 {
		c.AbortWithStatusJSON(400, gin.H{"error": "User already exists"})
		return
	}
	salt := MakeSalt(24)
	user.Salt = salt

	secret := os.Getenv("SECRET_KEY")
	fmt.Println("PASS + SALT", user.Password+salt)
	pass := Encrypt(user.Password+salt, secret)
	user.Password = pass

	res, err := GetDb().NamedExec(CreateUser, user)
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
		return
	}
	id, _ := res.LastInsertId()
	user.ID = int32(id)

	token, err := CreateJwt(user)

	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"jwt": token})
}

// LoginHandler authenticates a user and returns a JWT token
func LoginHandler(c *gin.Context) {
	type userLogin struct {
		Email    string `json:"Email" binding:"required"`
		Password string `json:"Password" binding:"required"`
	}
	var user userLogin
	err := c.BindJSON(&user)
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
		return
	}

	query := fmt.Sprintf(GetUserByEmail, user.Email)
	var results []User
	GetDb().Select(&results, query)

	if len(results) == 0 {
		c.AbortWithStatusJSON(400, gin.H{"error": "User does not exist"})
		return
	}
	secret := os.Getenv("SECRET_KEY")
	fmt.Println("PASS + SALT", user.Password+results[0].Salt)
	auth := CompareEncoded(user.Password+results[0].Salt, results[0].Password, secret)

	if auth == false {
		c.AbortWithStatusJSON(400, gin.H{"error": "Password is incorrect"})
		return
	}

	token, _ := CreateJwt(results[0])
	c.JSON(200, gin.H{"jwt": token})
}
