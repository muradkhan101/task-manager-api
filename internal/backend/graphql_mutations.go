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
		board.TaskOrder = "[]"
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
		paramsToStruct(&params, "board", &board)

		_, err := GetDb().NamedExec(UpdateBoard, board)
		return board, err
	},
}

var RemoveBoardMutation = &graphql.Field{
	Type:        graphql.Boolean,
	Description: "Remove a board",
	Args: graphql.FieldConfigArgument{
		"boardId": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.Int)},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		boardId, _ := params.Args["boardId"]
		queryStr := fmt.Sprintf(RemoveBoard, boardId)
		_, err := GetDb().Exec(queryStr)
		if err != nil {
			return false, err
		}
		return true, nil
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
		fmt.Println("[ISSUE]", issue.Status)
		return issue, err
	},
}

var RemoveIssueMutation = &graphql.Field{
	Type:        graphql.Boolean,
	Description: "Remove an issue",
	Args: graphql.FieldConfigArgument{
		"taskId": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.Int)},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		taskId, _ := params.Args["taskId"]
		queryStr := fmt.Sprintf(RemoveIssue, taskId)
		_, err := GetDb().Exec(queryStr)
		if err != nil {
			return false, err
		}
		return true, nil
	},
}

var UpdateTaskOrderMutation = &graphql.Field{
	Type:        BoardType,
	Description: "Update the order of tasks",
	Args: graphql.FieldConfigArgument{
		"TaskOrder": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
		"BoardId":   &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.Int)},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		taskOrder, _ := params.Args["TaskOrder"]
		boardId, _ := params.Args["BoardId"]
		queryStr := fmt.Sprintf(UpdateTaskOrder, taskOrder, boardId)
		_, err := GetDb().Exec(queryStr)
		if err != nil {
			return nil, err
		}

		var queryResult []Board
		query := fmt.Sprintf(GetBoardByID, boardId)
		err = GetDb().Select(&queryResult, query)
		return queryResult[0], err
	},
}

var UpdateBoardOrderMutation = &graphql.Field{
	Type:        UserType,
	Description: "Update the order of boards",
	Args: graphql.FieldConfigArgument{
		"BoardOrder": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
		"UserId":     &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.Int)},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		boardOrder, _ := params.Args["BoardOrder"]
		userId, _ := params.Args["UserId"]
		queryStr := fmt.Sprintf(UpdateBoardOrder, boardOrder, userId)
		_, err := GetDb().Exec(queryStr)
		if err != nil {
			return nil, err
		}

		var queryResult []User
		query := fmt.Sprintf(GetUserById, userId)
		err = GetDb().Select(&queryResult, query)
		return queryResult[0], err
	},
}
