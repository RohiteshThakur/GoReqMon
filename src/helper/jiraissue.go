package helper

import (
	"fmt"
	jira "github.com/andygrunwald/go-jira"
)

func JiraIssue(client *jira.Client) {
	// Search Issues based on JQL:
	searchoptions := &jira.SearchOptions{
		StartAt:		0,
		MaxResults:		50,
		Expand:			"",
		ValidateQuery: 	"validateQuery",
	}

	issues, _, i_err := client.Issue.Search("Area = \"<Area name>\" AND status = Open", searchoptions)		// update: Returns an array of Issue.
	if issues == nil {
		fmt.Println("No Issues found")
	}
	if i_err != nil {
		fmt.Printf("Error given: %s", i_err)
	} else {
		for _, elem := range issues {										// Iterate over issues.					
			print("\n\n----------------------------------------\n")
			// -- Print the whole thing.
			//PrettyPrint(elem)

			// -- Get Jira Issue Number
			fmt.Println("Issue#: ", elem.Key)
			//Get
			//PrettyPrint(elem.Fields.Status)								// These fields are defined in the respective source code.
			//print("\n\n")

			// -- Get Area - "<Are Name on Jira board>"
			//PrettyPrint(elem.Fields.Unknowns["customfield_<XXXXXX>"])		// "customfield_10109" is a dict key. - works!
			area := elem.Fields.Unknowns["customfield_<XXXXXX>"]				// value of "customfield_10109" is a dict and therefore requires assertion before we can access value.
			//fmt.Println(area)												// Go treats area as a map[] but doesn't knows type of its value so we need to assert.
			areav, _ := area.(map[string]interface{})						// Go needs an explicit assertion that what follows after customfield_* is a interface. 
			fmt.Println(areav["value"])

			// Get Task Type
			tt := elem.Fields.Unknowns["customfield_<XXXXXX>"]
			//fmt.Println(area)
			ttv, _ := tt.(map[string]interface{})						
			fmt.Println(ttv["value"])

			// Get requester ID
			flid := elem.Fields.Unknowns["customfield_<XXXXXX>"]
			fmt.Println (flid)

			// Get requester email ID
			emailid := elem.Fields.Unknowns["customfield_<XXXXXX>"]
			fmt.Println (emailid)

			// Get LM
			lm := elem.Fields.Unknowns["customfield_<XXXXXX>"]
			fmt.Println (lm)

			// Get VS Name
			vs := elem.Fields.Unknowns["customfield_<XXXXXX>"]
			//fmt.Println(area)
			vsv, _ := vs.(map[string]interface{})						
			fmt.Println(vsv["value"])

			// Get LN
			lab := elem.Fields.Unknowns["customfield_<XXXXXX>"]
			fmt.Println (lab)


		}
	}
}	
