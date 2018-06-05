package main

import "github.com/graphql-go/graphql"

// UserType is graphql object for basic user data
var UserType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "User",
		Fields: graphql.Fields{
			"ID":        &graphql.Field{Type: graphql.Int},
			"FirstName": &graphql.Field{Type: graphql.String},
			"LastName":  &graphql.Field{Type: graphql.String},
			"Email":     &graphql.Field{Type: graphql.String},
			"Password":  &graphql.Field{Type: graphql.String},
			"Salt":      &graphql.Field{Type: graphql.String},
		},
	},
)

// IssueType is a graphql object for an Issue
var IssueType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Issue",
		Fields: graphql.Fields{
			"ID":          &graphql.Field{Type: graphql.Int},
			"CreatedBy":   &graphql.Field{Type: graphql.Int},
			"CreatedDate": &graphql.Field{Type: graphql.String},
			"Owner":       &graphql.Field{Type: graphql.Int},
			"Name":        &graphql.Field{Type: graphql.String},
			"Description": &graphql.Field{Type: graphql.String},
			"Reporter":    &graphql.Field{Type: graphql.Int},
			"Updates":     &graphql.Field{Type: graphql.Int},
			"DueDate":     &graphql.Field{Type: graphql.String},
		},
	},
)

// BoardType is a graphql object for a board
var BoardType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Board",
		Fields: graphql.Fields{
			"ID":          &graphql.Field{Type: graphql.Int},
			"CreatedBy":   &graphql.Field{Type: graphql.Int},
			"CreatedDate": &graphql.Field{Type: graphql.String},
			"Owner":       &graphql.Field{Type: graphql.Int},
			"Name":        &graphql.Field{Type: graphql.String},
		},
	},
)
