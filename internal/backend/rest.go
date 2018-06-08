package backend

import (
	"errors"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	. "github.com/task-manager-api/internal/types"
)

type jwtToken struct {
	jwt string
}

// CreateUserHandler handles creation of a new user
func CreateUserHandler(c *gin.Context) {
	var user User
	c.BindJSON(&user)

	query := fmt.Sprintf(GetUserByEmail, user.Email)
	var results []User
	err := GetDb().Select(&results, query)

	if len(results) > 0 {
		c.AbortWithError(400, errors.New("User already exists"))
	}
	salt := MakeSalt(24)
	user.Salt = salt

	secret := os.Getenv("SECRET_KEY")
	pass := Encrypt(user.Password+salt, secret)
	user.Password = pass

	res, err := GetDb().NamedExec(CreateUser, user)
	if err != nil {
		c.AbortWithError(400, err)
	}
	id, _ := res.LastInsertId()
	user.ID = int32(id)

	token, err := CreateJwt(user)

	if err != nil {
		c.AbortWithError(400, err)
	}

	c.JSON(200, jwtToken{token})
}

// LoginHandler authenticates a user and returns a JWT token
func LoginHandler(c *gin.Context) {
	var user User
	c.BindJSON(&user)

	query := fmt.Sprintf(GetUserByEmail, user.Email)
	var results []User
	err := GetDb().Select(&results, query)

	if len(results) == 0 {
		c.AbortWithError(400, errors.New("User does not exist"))
	}
	pass := user.Password

	auth := CompareEncoded(user.Password+user.Salt, results[0].Password)

	if auth == false {
		c.AbortWithError(400, errors.New("Password is incorrect"))
	}

	token, _ := CreateJwt(user)
	c.JSON(200, jwtToken{token})
}
