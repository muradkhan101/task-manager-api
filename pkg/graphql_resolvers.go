package main

import (
	"github.com/graphql-go/graphql"
)

// IssueResolver retrieves list of Issues for a user / org
var IssueResolver = graphql.Fields{
	"issue": &graphql.Field{
		Type:        IssueType,
		Description: "Get issues for a user / organization by querying off ID",
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
		},
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			id, isOk := params.Args["id"]
			var queryResult []Issue
			if isOk {
				err := DB.Select(&queryResult, GetIssuesByOwner, id)
				if err != nil {
					return queryResult, err
				}
			}
			return queryResult, nil
		},
	},
}

// BoardResolver retrieves list of Boards for a user / org
var BoardResolver = graphql.Fields{
	"Board": &graphql.Field{
		Type:        BoardType,
		Description: "Get boards for a user / organization by querying off ID",
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
		},
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			id, isOk := params.Args["id"]
			var queryResult []Board
			if isOk {
				err := DB.Select(&queryResult, GetBoardsByOwner, id)
				if err != nil {
					return queryResult, err
				}
			}
			return queryResult, nil
		},
	},
}
