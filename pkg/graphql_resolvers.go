package main

import (
	"fmt"

	"github.com/graphql-go/graphql"
)

// UserResolver retreieves user info based off ID
var userResolver = &graphql.Field{
	Type:        IssueType,
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
			err := DB.Select(&queryResult, GetUserInfo, id)
			if err != nil {
				return queryResult, err
			}
		}
		return queryResult, nil
	},
}

// BoardResolver retrieves info about a board based off ID
var boardResolver = &graphql.Field{
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
			err := DB.Select(&queryResult, GetBoardByID, id)
			if err != nil {
				return queryResult, err
			}
		}
		return queryResult, nil
	},
}

var rootQuery = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "RootQuery",
		Fields: &graphql.Fields{
			"user":  userResolver,
			"board": boardResolver,
		},
	},
)

var schema, _ = graphql.NewSchema(
	graphql.SchemaConfig{
		Query: rootQuery,
	},
)

func queryGraphql(query string, schema graphql.Schema) *graphql.Result {
	result := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: query,
	})
	if len(result.Errors) > 0 {
		fmt.Printf("Wrong result, errors occured: %v", result.Errors)
	}
	return result
}

// ExecuteQuery executes the given graphql query and returns the result
func ExecuteQuery(query string) *graphql.Result {
	return queryGraphql(query, schema)
}
