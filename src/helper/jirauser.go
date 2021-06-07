package helper

import (
	"fmt"
	jira "github.com/andygrunwald/go-jira"
)

func JiraUser(client *jira.Client) {
	user, _, err := client.User.Get("<user_name>")								// update:
	if err != nil {
		fmt.Printf("\nerror: %v\n", err)
		return 
	} else {
		fmt.Printf("\nEmail: %v\nSuccess!\n", user.EmailAddress)
	}	
}