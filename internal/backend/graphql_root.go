package backend

import (
	"fmt"

	"github.com/graphql-go/graphql"
)

var rootQuery = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "RootQuery",
		Fields: graphql.Fields{
			"user":  UserResolver,
			"board": BoardResolver,
		},
	},
)

var rootMutation = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "RootMutation",
		Fields: graphql.Fields{
			"createBoard":      CreateBoardMutation,
			"updateBoard":      UpdateBoardMutation,
			"createIssue":      CreateIssueMutation,
			"updateIssue":      UpdateIssueMutation,
			"updateTaskOrder":  UpdateTaskOrderMutation,
			"updateBoardOrder": UpdateBoardOrderMutation,
		},
	},
)

var schema, _ = graphql.NewSchema(
	graphql.SchemaConfig{
		Query:    rootQuery,
		Mutation: rootMutation,
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
