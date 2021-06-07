package helper

import (
	"fmt"
	jira "github.com/andygrunwald/go-jira"
)

var boardsListOptions *jira.BoardListOptions

func JiraBoard(client *jira.Client) {

	boardsListOptions := &jira.BoardListOptions{
		BoardType:      "",
		Name:           "<JIRA BOARD NAME>",              // update:
		ProjectKeyOrID: "",
	}
	boardsListOptions.StartAt = 1
	boardsListOptions.MaxResults = 100

	
	boards, _, err := client.Board.GetAllBoards(boardsListOptions)
	if boards == nil {
		fmt.Println("Expected boards list. Boards list is nil")
		fmt.Println("Check whether user is connected to Sandbox")
	}
	if err != nil {
		fmt.Printf("Error given: %s", err)
	} else {
		fmt.Println("\nBoards:")
		fmt.Println(boards)
	}
}	