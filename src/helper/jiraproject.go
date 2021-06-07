package helper

import (
	"fmt"
	"reflect"
	jira "github.com/andygrunwald/go-jira"

)


func JiraProject(client *jira.Client){
	projects, _, p_err := client.Project.GetList()
	if projects == nil {
		fmt.Println("Expected boards list. Boards list is nil")
	}
	if p_err != nil {
		fmt.Printf("Error given: %s", p_err)
	} else {
		fmt.Println("success")
		//fmt.Printf("%+v\n", projects)
		//PrettyPrint(*projects) //Or PrettyPrint(projects) // same result.
		fmt.Println(reflect.TypeOf(projects))   			// returns: *jira.ProjectList (a pointer). Therefore, *projects is jira.ProjectList, which is an array of struct.
		for _, elem := range *projects {					// So we iterate over, *projects (array of structs)
			if (elem.Name == "<Jira Board/Project name>"){
				fmt.Println(elem)
			}
			fmt.Println()
		}

		/* Test Marshalling - 
		byte_array, bm_err := json.Marshal(projects)
		if (bm_err != nil){
			panic(bm_err)
		} else {
			PrettyPrint(string(byte_array))
		}
		*/		
	}
}

func JiraFilterProject(client *jira.Client) {
	filter_projects, _, err := client.Project.ListWithOptions(&jira.GetQueryOptions{Expand: "name"})
	if filter_projects == nil {
		fmt.Println("No projects found.")
	}
	if err != nil {
		fmt.Printf("Error given: %s", err)
	} else {
		PrettyPrint(filter_projects)
	}
}