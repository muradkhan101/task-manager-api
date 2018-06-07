package main

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestBoardJson(t *testing.T) {
	var boardJson Board
	jsonString := `{"ID":1,"CreatedBy":22,"CreatedDate":"Today","Owner":22,"Name":"Board One"}`
	err := json.Unmarshal([]byte(jsonString), &boardJson)
	// issues := make([]Issue, 0)
	// boardJson := Board{1, 22, "Today", 22, "Board One", issues}
	// jsonStr, err := json.Marshal(&boardJson)
	if err != nil {
		fmt.Print(err)
		t.Errorf("Failed to decode JSON")
	}
}

func TestUserJson(t *testing.T) {
	var userJson User
	jsonString := `{"ID":12, "FirstName":"Potato", "LastName":"Header", "Email":"k@g.com", "Password": "123456", "Salt":"frechfries"}`
	err := json.Unmarshal([]byte(jsonString), &userJson)
	// user := User{111, "Bill", "Murray", "kop@gm.com", "123456", "saltypoo"}
	// bytes, err := json.Marshal(&user)
	if err != nil {
		fmt.Print(err)
		t.Errorf("Failed to unmarshal JSON")
	} else {
		fmt.Print(jsonString)
	}
}

func TestIssueJson(t *testing.T) {
	var issueJson Issue
	jsonStr := `{
		"ID":123,
		"CreatedBy": 456,
		"CreatedDate": "12/21/1990",
		"Owner": 6969,
		"Name": "Boil potatoes",
		"Description": "Boil them, mash them, put them in a stew",
		"Reporter": 1111,
		"Updates": 11,
		"DueDate": "Tomorrow"
		}`
	err := json.Unmarshal([]byte(jsonStr), &issueJson)
	if err != nil {
		fmt.Print(err)
		t.Errorf("Failed to unmarshal JSON")
	}
}
