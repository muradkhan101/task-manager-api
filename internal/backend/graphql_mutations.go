package backend

import (
	"fmt"

	"encoding/json"

	"github.com/graphql-go/graphql"
	. "github.com/task-manager-api/internal/types"
)

func paramsToStruct(params *graphql.ResolveParams, param string, item interface{}) error {
	paramString, _ := params.Args[param]
	data, err := json.Marshal(paramString)
	if err != nil {
		return err
	}
	err = json.Unmarshal([]byte(data), item)
	return err

}

// CreateBoardMutation creates a new board in the DB and returns it
var CreateBoardMutation = &graphql.Field{
	Type:        BoardType,
	Description: "Create a new board",
	Args: graphql.FieldConfigArgument{
		"board": &graphql.ArgumentConfig{Type: BoardInput},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		var board Board
		err := paramsToStruct(&params, "board", &board)

		res, err := GetDb().NamedExec(CreateBoard, board)

		id, _ := res.LastInsertId()
		board.ID = int32(id)
		fmt.Println(board)
		return board, err
	},
}

// UpdateBoardMutation updates a board existing in the DB
var UpdateBoardMutation = &graphql.Field{
	Type:        BoardType,
	Description: "Update an existing board",
	Args: graphql.FieldConfigArgument{
		"board": &graphql.ArgumentConfig{Type: BoardInput},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		var board Board
		err := paramsToStruct(&params, "board", &board)

		_, err := GetDb().NamedExec(UpdateBoard, board)
		return board, err
	},
}

// CreateIssueMutation creates a new issue in the DB and returns it
var CreateIssueMutation = &graphql.Field{
	Type:        IssueType,
	Description: "Create a new issue",
	Args: graphql.FieldConfigArgument{
		"issue": &graphql.ArgumentConfig{Type: IssueInput},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		var issue Issue
		err := paramsToStruct(&params, "issue", &issue)
		res, err := GetDb().NamedExec(CreateIssue, issue)

		id, _ := res.LastInsertId()
		issue.ID = int32(id)
		return issue, err
	},
}

// UpdateIssueMutation updates an issue existing in the DB
var UpdateIssueMutation = &graphql.Field{
	Type:        IssueType,
	Description: "Update an issue",
	Args: graphql.FieldConfigArgument{
		"issue": &graphql.ArgumentConfig{Type: IssueInput},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		var issue Issue
		paramsToStruct(&params, "issue", &issue)
		_, err := GetDb().NamedExec(UpdateIssue, issue)

		return issue, err
	},
}
