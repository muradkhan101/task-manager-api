package main

import (
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/service/rds/rdsutils"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

const (
	endpoint = "taskmaster.c1vghsmtnzjg.us-west-2.rds.amazonaws.com"
	region   = "us-west-2"
	dbUser   = "mkhan"
	dbName   = "main"
)

const (
	GetIssuesByOwner = "SELECT * FROM issues WHERE owner = %d"
	GettIssueById    = "SELECT * FROM issues WHERE id = %d"
	GetIssuesByBoard = "SELECT * FROM issues WHERE board = %d"
	GetBoardByID     = "SELECT * FROM boards WHERE id = %d"
	GetBoardsByOwner = "SELECT * FROM boards WHERE owner = %d"
	GetUserInfo      = "SELECT id, first_name, last_name, email FROM users where id = %d"
)

// Connect s to AWS RDS instance using credentials in environment variables
func Connect() *sqlx.DB {
	fmt.Print("test")
	creds := credentials.NewEnvCredentials()
	authToken, err := rdsutils.BuildAuthToken(endpoint, region, dbUser, creds)
	if err != nil {
		log.Fatal("Failed to load AWS credentials")
	}
	connectStr := fmt.Sprintf("%s:%s@tcp(%s)/%s?tls=true", dbUser, authToken, endpoint, dbName)
	db, err := sqlx.Connect("mysql", connectStr)
	if err != nil {
		log.Fatal("Failed to connect to DB")
	}
	return db
}
