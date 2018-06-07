package main

import (
	"fmt"

	"github.com/graphql-go/graphql"
)

// UserType is graphql object for user data with ability to get ISsues and Boards
var UserType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "User",
		Fields: graphql.Fields{
			"ID":        &graphql.Field{Type: graphql.NewNonNull(graphql.Int)},
			"FirstName": &graphql.Field{Type: graphql.NewNonNull(graphql.String)},
			"LastName":  &graphql.Field{Type: graphql.NewNonNull(graphql.String)},
			"Email":     &graphql.Field{Type: graphql.NewNonNull(graphql.String)},
			"Issues": &graphql.Field{
				Type:        graphql.NewList(IssueType),
				Description: "Get issues for a user by querying off ID",
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					var queryResult []Issue
					user, isOk := params.Source.(User)
					if isOk {
						query := fmt.Sprintf(GetIssuesByOwner, user.ID)
						fmt.Println("query: ", query)
						err := DB.Select(&queryResult, query)
						return queryResult, err
					}
					issue := make([]Issue, 1)
					issue[0] = Issue{123, 456, "12/21/1990", 6969, "Boil potatoes", "Boil them, mash them, put them in a stew", 1111, 11, "Tomorrow", 666}
					return issue, nil
				},
			},
			"Boards": &graphql.Field{
				Type:        graphql.NewList(BoardType),
				Description: "Get boards for a user by querying off ID",
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					var queryResult []Board
					user, isOk := params.Source.(User)
					if isOk {
						query := fmt.Sprintf(GetBoardsByOwner, user.ID)
						fmt.Println("query: ", query)
						err := DB.Select(&queryResult, query)
						return queryResult, err
					}
					issue := make([]Issue, 1)
					issue[0] = Issue{123, 456, "12/21/1990", 6969, "Boil potatoes", "Boil them, mash them, put them in a stew", 1111, 11, "Tomorrow", 666}
					board := make([]Board, 1)
					board[0] = Board{123, 456, "12/21/1990", 6969, "Boil potatoes", 1, issue}
					return board, nil
				},
			},
		},
	},
)

// IssueType is a graphql object for an Issue
var IssueType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Issue",
		Fields: graphql.Fields{
			"ID":          &graphql.Field{Type: graphql.NewNonNull(graphql.Int)},
			"CreatedBy":   &graphql.Field{Type: graphql.NewNonNull(graphql.Int)},
			"CreateDate":  &graphql.Field{Type: graphql.NewNonNull(graphql.String)},
			"Owner":       &graphql.Field{Type: graphql.NewNonNull(graphql.Int)},
			"Name":        &graphql.Field{Type: graphql.NewNonNull(graphql.String)},
			"Description": &graphql.Field{Type: graphql.NewNonNull(graphql.String)},
			"Reporter":    &graphql.Field{Type: graphql.Int},
			"Updates":     &graphql.Field{Type: graphql.NewNonNull(graphql.Int)},
			"DueDate":     &graphql.Field{Type: graphql.String},
			"Board":       &graphql.Field{Type: graphql.NewNonNull(graphql.Int)},
		},
	},
)

// BoardType is a graphql object for a board with option to get issues
var BoardType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Board",
		Fields: graphql.Fields{
			"ID":         &graphql.Field{Type: graphql.NewNonNull(graphql.Int)},
			"CreatedBy":  &graphql.Field{Type: graphql.NewNonNull(graphql.Int)},
			"CreateDate": &graphql.Field{Type: graphql.NewNonNull(graphql.String)},
			"Owner":      &graphql.Field{Type: graphql.NewNonNull(graphql.Int)},
			"Name":       &graphql.Field{Type: graphql.NewNonNull(graphql.String)},
			"Issues": &graphql.Field{
				Type:        graphql.NewList(IssueType),
				Description: "Get list of issues on a board",
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					var queryResult []Issue
					board, isOk := params.Source.(Board)
					if isOk {
						query := fmt.Sprintf(GetIssuesByBoard, board.ID)
						fmt.Println("query: ", query)
						err := DB.Select(&queryResult, query)
						return queryResult, err
					}
					issue := make([]Issue, 1)
					issue[0] = Issue{123, 456, "12/21/1990", 6969, "Boil potatoes", "Boil them, mash them, put them in a stew", 1111, 11, "Tomorrow", 666}
					return issue, nil
				},
			},
		},
	},
)
