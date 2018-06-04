package main

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestBoardJson(t *testing.T) {
	var boardJson Board
	jsonString := `{"ID": 1,"CreatedBy": 22,"CreatedDate": "12/15/2014","Owner": 41,"Name": "TestBoard", "Issues": []}`
	err := json.Unmarshal([]byte(jsonString), &boardJson)
	if err == nil {
		fmt.Print(boardJson)
	} else {
		t.Errorf("Failed to decode JSON")
	}
}

func TestUserJson(t *testing.T) {
	var userJson User
	jsonString := `{"ID":12,"FirstName":"Potato","LastName":"Header","Email":"k@g.com",Password":"123456","Salt":"frechfries"}`
	err := json.Unmarshal([]byte(jsonString), &jsonString)
	if err != nil {
		fmt.Print(err)
		t.Errorf("Failed to unmarshal JSON")
	} else {
		fmt.Print(userJson)
	}
}
