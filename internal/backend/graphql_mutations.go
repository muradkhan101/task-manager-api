package backend

import (
	"fmt"

	"encoding/json"

	"github.com/graphql-go/graphql"
	. "github.com/task-manager-api/internal/types"
)

// CreateBoard creates a new board in the DB and returns it
var CreateBoard = &graphql.Field{
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
		fmt.Println("BOARD:  ", board)
		// DB.Exec()
		return board, nil
	},
}

// UpdateBoard udpates a board existing in the DB
var UpdateBoard = &graphql.Field{
	Type:        BoardType,
	Description: "Update an existing board",
	Args: graphql.FieldConfigArgument{
		"board": &graphql.ArgumentConfig{Type: BoardInput},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		board, _ := params.Args["board"].(Board)
		// DB.Exec()
		return board, nil
	},
}
