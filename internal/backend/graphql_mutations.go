package backend

import (
	"fmt"

	"encoding/json"

	"github.com/graphql-go/graphql"
	. "github.com/task-manager-api/internal/types"
)

// CreateBoardMutation creates a new board in the DB and returns it
var CreateBoardMutation = &graphql.Field{
	Type:        BoardType,
	Description: "Create a new board",
	Args: graphql.FieldConfigArgument{
		"board": &graphql.ArgumentConfig{Type: BoardInput},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		boardString, _ := params.Args["board"]
		data, _ := json.Marshal(boardString)

		var board Board
		json.Unmarshal([]byte(data), &board)

		res, err := GetDb().NamedExec(CreateBoard, board)

		id, _ := res.LastInsertId()
		board.ID = int32(id)
		fmt.Println(board)
		return board, err
	},
}

// UpdateBoardMutation udpates a board existing in the DB
var UpdateBoardMutation = &graphql.Field{
	Type:        BoardType,
	Description: "Update an existing board",
	Args: graphql.FieldConfigArgument{
		"board": &graphql.ArgumentConfig{Type: BoardInput},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		boardString, _ := params.Args["board"]
		data, _ := json.Marshal(boardString)

		var board Board
		json.Unmarshal([]byte(data), &board)

		res, err := GetDb().NamedExec(UpdateBoard, board)
		return board, err
	},
}
