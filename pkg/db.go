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
	GetIssuesByOwner = "SELECT * FROM issues WHERE owner = %s"
	GettIssueById    = "SELECT * FROM issues WHERE id = %s"
	GetBoardsByOwner = "SELECT * FROM boards WHERE owner = %s"
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
