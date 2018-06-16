package types

import (
	"database/sql"
	"time"

	"github.com/go-sql-driver/mysql"
)

// Issue contains information about a specific task to be worked on by the user
type Issue struct {
	ID          int32         `json:"ID" db:"id" binding:"required"`
	CreatedBy   int32         `json:"CreatedBy" db:"created_by" binding:"required"`
	CreateDate  time.Time     `json:"CreateDate" db:"create_date" binding:"required"`
	Owner       int32         `json:"Owner" db:"owner" binding:"required"`
	Name        string        `json:"Name" db:"name" binding:"required"`
	Description string        `json:"Description" db:"description" binding:"required"`
	Reporter    sql.NullInt64 `json:"Reporter,omitempty" db:"reporter"`
	// Updates     int32  `json:"Updates" db:"updates"`
	DueDate mysql.NullTime `json:"DueDate,omitempty" db:"due_date"`
	Board   int32          `json:"Board" db:"board" binding:"required"`
	Status  int32          `json:"Status" db:"status" binding:"required"`
}

// User is a user object used for authentication
type User struct {
	ID         int32   `json:"ID" db:"id" binding:"required"`
	FirstName  string  `json:"FirstName" db:"first_name" binding:"required"`
	LastName   string  `json:"LastName" db:"last_name" binding:"required"`
	Email      string  `json:"Email" db:"email" binding:"required"`
	Password   string  `json:"Password" db:"password"`
	Salt       string  `json:"Salt" db:"salt"`
	Issues     []Issue `json:"Issues"`
	Boards     []Board `json:"Boards"`
	BoardOrder string  `json:"BoardOrder,omitempty" db:"board_order"`
}

// Board contains a list of issues and info about board
type Board struct {
	ID         int32     `json:"ID" db:"id" binding:"required"`
	CreatedBy  int32     `json:"CreatedBy" db:"created_by" binding:"required"`
	CreateDate time.Time `json:"CreateDate" db:"create_date" binding:"required"`
	Owner      int32     `json:"Owner" db:"owner" binding:"required"`
	Name       string    `json:"Name" db:"name" binding:"required"`
	// Updates    int32   `json:"Updates" db:"updates"`
	Issues     []Issue `json:"Issues"`
	IssueOrder string  `json:"IssueOrder,omitempty" db:"issue_order"`
}
