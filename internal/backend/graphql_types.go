package backend

import (
	"fmt"

	"github.com/graphql-go/graphql"
	. "github.com/task-manager-api/internal/types"
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
						err := GetDb().Select(&queryResult, query)
						return queryResult, err
					}
					return queryResult, nil
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
						err := GetDb().Select(&queryResult, query)
						return queryResult, err
					}
					return queryResult, nil
				},
			},
			"BoardOrder": &graphql.Field{Type: graphql.String},
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
			"CreateDate":  &graphql.Field{Type: graphql.NewNonNull(graphql.DateTime)},
			"Owner":       &graphql.Field{Type: graphql.NewNonNull(graphql.Int)},
			"Name":        &graphql.Field{Type: graphql.NewNonNull(graphql.String)},
			"Description": &graphql.Field{Type: graphql.NewNonNull(graphql.String)},
			"Reporter":    &graphql.Field{Type: graphql.Int},
			// "Updates":     &graphql.Field{Type: graphql.NewNonNull(graphql.Int)},
			"DueDate": &graphql.Field{Type: graphql.DateTime},
			"Board":   &graphql.Field{Type: graphql.NewNonNull(graphql.Int)},
			"Status":  &graphql.Field{Type: graphql.NewNonNull(graphql.Int)},
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
			"CreateDate": &graphql.Field{Type: graphql.NewNonNull(graphql.DateTime)},
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
						err := GetDb().Select(&queryResult, query)
						return queryResult, err
					}
					return queryResult, nil
				},
			},
			"TaskOrder": &graphql.Field{Type: graphql.String},
		},
	},
)

// BoardInput is definition for board mutations
var BoardInput = graphql.NewInputObject(
	graphql.InputObjectConfig{
		Name: "BoardInput",
		Fields: graphql.InputObjectConfigFieldMap{
			"ID":         &graphql.InputObjectFieldConfig{Type: graphql.Int},
			"CreatedBy":  &graphql.InputObjectFieldConfig{Type: graphql.NewNonNull(graphql.Int)},
			"CreateDate": &graphql.InputObjectFieldConfig{Type: graphql.NewNonNull(graphql.DateTime)},
			"Owner":      &graphql.InputObjectFieldConfig{Type: graphql.NewNonNull(graphql.Int)},
			"Name":       &graphql.InputObjectFieldConfig{Type: graphql.NewNonNull(graphql.String)},
		},
	},
)

// IssueInput is definiton for Issue mutations
var IssueInput = graphql.NewInputObject(
	graphql.InputObjectConfig{
		Name: "IssueInput",
		Fields: graphql.InputObjectConfigFieldMap{
			"ID":          &graphql.InputObjectFieldConfig{Type: graphql.Int},
			"CreatedBy":   &graphql.InputObjectFieldConfig{Type: graphql.NewNonNull(graphql.Int)},
			"CreateDate":  &graphql.InputObjectFieldConfig{Type: graphql.NewNonNull(graphql.DateTime)},
			"Owner":       &graphql.InputObjectFieldConfig{Type: graphql.NewNonNull(graphql.Int)},
			"Name":        &graphql.InputObjectFieldConfig{Type: graphql.NewNonNull(graphql.String)},
			"Description": &graphql.InputObjectFieldConfig{Type: graphql.NewNonNull(graphql.String)},
			"Reporter":    &graphql.InputObjectFieldConfig{Type: graphql.Int},
			"DueDate":     &graphql.InputObjectFieldConfig{Type: graphql.DateTime},
			"Board":       &graphql.InputObjectFieldConfig{Type: graphql.NewNonNull(graphql.Int)},
			"Status":      &graphql.InputObjectFieldConfig{Type: graphql.NewNonNull(graphql.Int)},
		},
	},
)

// UserInput is definiton for User mutations
var UserInput = graphql.NewInputObject(
	graphql.InputObjectConfig{
		Name: "IssueInput",
		Fields: graphql.InputObjectConfigFieldMap{
			"ID":        &graphql.InputObjectFieldConfig{Type: graphql.Int},
			"FirstName": &graphql.InputObjectFieldConfig{Type: graphql.NewNonNull(graphql.String)},
			"LastName":  &graphql.InputObjectFieldConfig{Type: graphql.NewNonNull(graphql.String)},
			"Email":     &graphql.InputObjectFieldConfig{Type: graphql.NewNonNull(graphql.String)},
			"Password":  &graphql.InputObjectFieldConfig{Type: graphql.NewNonNull(graphql.String)},
			"Salt":      &graphql.InputObjectFieldConfig{Type: graphql.NewNonNull(graphql.String)},
		},
	},
)
