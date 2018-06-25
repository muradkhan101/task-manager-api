package backend

import (
	"fmt"

	"github.com/graphql-go/graphql"
	. "github.com/task-manager-api/internal/types"
)

// UserResolver retreieves user info based off ID
var UserResolver = &graphql.Field{
	Type:        UserType,
	Description: "Get info about a user",
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		id, isOk := params.Args["id"]
		var queryResult User
		if isOk {
			query := fmt.Sprintf(GetUserById, id)
			row := GetDb().QueryRow(query)
			row.Scan(&queryResult)
			fmt.Println(queryResult)
			return queryResult, nil
		}
		return queryResult, nil
	},
}

// BoardResolver retrieves info about a board based off ID
var BoardResolver = &graphql.Field{
	Type:        BoardType,
	Description: "Get info about a board",
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		id, isOk := params.Args["id"]
		var queryResult Board
		if isOk {
			query := fmt.Sprintf(GetBoardByID, id)
			row := GetDb().QueryRow(query)
			row.Scan(&queryResult)
			fmt.Println(queryResult)
			return queryResult, nil
		}
		return queryResult, nil
	},
}
