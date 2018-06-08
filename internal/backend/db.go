package backend

import (
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

const (
	endpoint = "taskmaster.c1vghsmtnzjg.us-west-2.rds.amazonaws.com"
	region   = "us-west-2"
	dbUser   = "mkhan"
	dbName   = "main"
)

// SELECT statements for querying DB
const (
	GetIssuesByOwner = "SELECT * FROM issues WHERE owner = %d;"
	GettIssueById    = "SELECT * FROM issues WHERE id = %d;"
	GetIssuesByBoard = "SELECT * FROM issues WHERE board = %d;"
	GetBoardByID     = "SELECT * FROM boards WHERE id = %d;"
	GetBoardsByOwner = "SELECT * FROM boards WHERE owner = %d;"
	GetUserInfo      = "SELECT id, first_name, last_name, email FROM users where id = %d;"
)

// CREATE and UPDATE statements for db
const (
	CreateBoard = "INSERT INTO boards (name, created_by, create_date, owner) VALUES (:name, :created_by, :create_date, :owner);"
	UpdateBoard = "UPDATE boards SET name = :name WHERE id = :id;"
)

func setUpDb() func() *sqlx.DB {
	var db *sqlx.DB
	return func() *sqlx.DB {
		if db != nil {
			return db
		}
		password := os.Getenv("RDS_PASSWORD")
		connectStr := fmt.Sprintf("%s:%s@tcp(%s)/%s", dbUser, password, endpoint, dbName)
		db, err := sqlx.Connect("mysql", connectStr)
		if err != nil {
			log.Fatal("Failed to connect to DB")
		}
		return db
	}
}

// GetDb returns a AWS RDS instance using credentials in environment variables
var GetDb = setUpDb()
