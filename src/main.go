package main

import (
	//"os"
	//"fmt"
	//"sync"
	"azure"
	"helper"

	"context"
	//"github.com/Azure/azure-sdk-for-go/services/aad/mgmt/2017-04-01/aad"
	"github.com/Azure/azure-sdk-for-go/services/graphrbac/1.6/graphrbac"
	"github.com/Azure/go-autorest/autorest/azure/auth"
	_"structutils"
	_"strings"

	"github.com/Azure/go-autorest/autorest/to"
)

func getADGroupsClient(fs auth.FileSettings) graphrbac.GroupsClient {
	groupsClient := graphrbac.NewGroupsClient(azure.GetTenantId(fs))
	a, _ := azure.GetGraphAuthorizer(fs)
	groupsClient.Authorizer = a
	groupsClient.AddToUserAgent(azure.UserAgent())
	return groupsClient
}

func CreateADGroup(ctx context.Context, fs auth.FileSettings) (graphrbac.ADGroup, error) {
	groupClient := getADGroupsClient(fs)
	return groupClient.Create(ctx, graphrbac.GroupCreateParameters{
		DisplayName:     to.StringPtr("GoSDKSamples"),
		MailEnabled:     to.BoolPtr(false),
		MailNickname:    to.StringPtr("GoSDKMN"),
		SecurityEnabled: to.BoolPtr(true),
	})
}

func getADUsersClient(fs auth.FileSettings) graphrbac.UsersClient {
	userClient := graphrbac.NewUsersClient(azure.GetTenantId(fs))
	a, _ := azure.GetGraphAuthorizer(fs)
	userClient.Authorizer = a
	//userClient.AddToUserAgent(azure.UserAgent())
	return userClient
}

func CreateADUser(ctx context.Context, fs auth.FileSettings) (graphrbac.User, error) {
	m := make(map[string]interface{})
	//m["forceChangePasswordNextSignInWithMfa"] = "false"
	// update: <password_to_assign> 
	passwordProfile := graphrbac.PasswordProfile{m, to.StringPtr("<password_to_assign>"), to.BoolPtr(false)}

	userClient := getADUsersClient(fs)
	return userClient.Create(ctx, graphrbac.UserCreateParameters {
		DisplayName:		to.StringPtr("gosampleuser"),
		UserPrincipalName:	to.StringPtr("gosampleuser@<tenant_name>"),                          /* update: tenant_name = jon.doe@micosoft.com */
		MailNickname:		to.StringPtr("GoSDKMN"),
		AccountEnabled:		to.BoolPtr(true),
		PasswordProfile:	&passwordProfile,
	})
}


func main (){
	client, _ := helper.JiraAuth()
	helper.JiraBoard(client)
	helper.JiraProject(client)
	helper.JiraIssue(client)

/*
	// Azure Auth
	os.Setenv("AZURE_AUTH_LOCATION", "/<path>/<to>/.azure/quickstart.auth")                     // update 
	fmt.Println(os.Getenv("AZURE_AUTH_LOCATION"))

	fs, _ := azure.MakeSettingsFile()

	/*
	sessAuth, err := azure.AzureAuthenticator(fs) 
	if err != nil {
		panic(err)
	}
	//fmt.Println(sessAuth.SubscriptionId)

	
	var wg sync.WaitGroup
	groups, err := azure.GetResourceGroups(sessAuth)
	if err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}

	fmt.Println ("Resource Groups: ", groups)


	wg.Add(len(groups))

	for _, group := range groups {
		go azure.GetVM(sessAuth, group, &wg)
	}
	wg.Wait()

	fmt.Println("\nGraphRBAC section:")
	graphrbacGroupsClient := getADGroupsClient(fs)

	//var body []byte
	for list, err := graphrbacGroupsClient.ListComplete(context.Background(), ""); list.NotDone(); err = list.Next() {
		if err != nil {
			fmt.Println("got error while traverising RG list: ", err)
		}
	
		adGroup := list.Value()
		fmt.Printf("%s\n", *adGroup.DisplayName)


		//structutils.ListType(GroupListResultIterator)
		//structutils.ListMethods(GroupListResultIterator)
		//structutils.ListFields(GroupListResultIterator)
	}



	newADGroup, err := CreateADGroup(context.Background(), fs)
	if (err != nil){
		fmt.Println("AD Group creation failed: ", err)
	} else {
		fmt.Println("\n New Group name: ", *newADGroup.DisplayName, "   ObjectType: ", newADGroup.ObjectType)
	}



    page, err := azure.GetADGroups(sessAuth) 
	fmt.Println(page.Values())



	opClient := aad.NewOperationsClient()
	opDiscoveryCollection, _ := opClient.List(context.Background())


	for _, elem := range opDiscoveryCollection {
		fmt.Println(elem)
	}
	fmt.Println(opDiscoveryCollection.Value)
	structutils.ListType(opDiscoveryCollection)




    // User List section
	adUserClient := getADUserClient(fs)  
	for list, err := adUserClient.ListComplete(context.Background(), ""); list.NotDone(); err = list.Next() {
		if err != nil {
			fmt.Print("got error while traverising User list: ", err)
		} 
		i := list.Value()
		fmt.Println("User Display Name:- ", *i.DisplayName, "  User Object ID:- ", *i.ObjectID)
		fmt.Println(i.UserType)
		fmt.Println(i.ObjectType)
	
	}


	newUser, err := CreateADUser(context.Background(), fs)
	if (err != nil){
		fmt.Println("AD User creation failed: ", err)
	} else {
		fmt.Println("\n New Group name: ", *newUser.DisplayName, "   ObjectType: ", newUser.ObjectType)
	}

*/
}