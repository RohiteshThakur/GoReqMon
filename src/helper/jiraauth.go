package helper

import (
	"fmt"
	"os"
	"encoding/json"
	"strings"
	jira "github.com/andygrunwald/go-jira"
)

type jiraAuthConfig struct {
	Jiraurl		string
	Username	string
	Password	string
}

func PrettyPrint(v interface{}) (err error) {
	b, err := json.MarshalIndent(v, "", "  ")
	if err == nil {
			fmt.Println(string(b))
	}
	return
}

func JiraAuth() (*jira.Client, error) {

	// Open jira_auth.json file.
	authConfigFile, err := os.Open("/<path>/<to>/.jira/jira_auth.json")                         // update: 
	if err != nil {
		fmt.Printf("\nCan't open config file: %v\n", err)
		return nil, err
	}
	defer authConfigFile.Close()

	conf := jiraAuthConfig{}
	decoder := json.NewDecoder(authConfigFile)
	decode_err := decoder.Decode(&conf)

	if (decode_err != nil) {
		fmt.Println("Couldn't Decode jiraAuth file.", decode_err)
		// invalid character 'X' after object key :- quotes(") missing in one of the keys.
	}

	fmt.Println(conf.Jiraurl)
	fmt.Println(conf.Username)


	tp := jira.BasicAuthTransport{
		Username: strings.TrimSpace(conf.Username),
		Password: strings.TrimSpace(conf.Password),
	}

	client, err := jira.NewClient(tp.Client(), strings.TrimSpace(conf.Jiraurl))
	if err != nil {
		fmt.Printf("\nerror: %v\n", err)
		return nil, err
	} else {
		return client, nil
	}
}



	


	


