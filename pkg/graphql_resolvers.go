package main

import (
	"fmt"

	"github.com/graphql-go/graphql"
)

// UserResolver retreieves user info based off ID
var userResolver = &graphql.Field{
	Type:        UserType,
	Description: "Get info about a user",
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		id, isOk := params.Args["id"]
		var queryResult []User
		if isOk {
			query := fmt.Sprintf(GetUserInfo, id)
			fmt.Println("query: ", query)
			err := DB.Select(&queryResult, query)
			return queryResult[0], err
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
		var queryResult []Board
		if isOk {
			query := fmt.Sprintf(GetBoardByID, id)
			fmt.Println("query: ", query)
			err := DB.Select(&queryResult, query)
			return queryResult[0], err
		}
		return queryResult, nil
	},
}

var rootQuery = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "RootQuery",
		Fields: graphql.Fields{
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
