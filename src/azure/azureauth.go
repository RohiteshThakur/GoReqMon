package azure

import (
	"fmt"
	//"structutils"
	"github.com/Azure/go-autorest/autorest"
	Azure "github.com/Azure/go-autorest/autorest/azure" 
	"github.com/Azure/go-autorest/autorest/azure/auth"
)

type AzureSession struct {
	SubscriptionId  string
	TenantId		string
	Authorizer      autorest.Authorizer
}

func MakeSettingsFile() (auth.FileSettings, error) {
	// Looks for env var: AZURE_AUTH_LOCATION" set to path of credentials file.
	// Credential file was created using: az ad sp create-for-rbac --sdk-auth > ~/.azure/quickstart.auth
	
	fs, err := auth.GetSettingsFromFile()

	//structutils.ListFields(fs)
	//structutils.ListMethods(fs)
	//structutils.ListType(fs)
	//fmt.Println(fs.Values)
	//fmt.Println(fs.Values["AZURE_SUBSCRIPTION_ID"])
	//fmt.Println(fs.Values["AZURE_TENANT_ID"])
	//print (fs.GetSubscriptionID())

	if (err != nil) {
		fmt.Println("Failed to create Settingsfile", err)
		panic(err)
	} else {
		return fs, nil
	}
}

func AzureAuthenticator (fs auth.FileSettings) (*AzureSession, error) {
	authorizer, err := auth.NewAuthorizerFromFile(Azure.PublicCloud.ResourceManagerEndpoint)  
	//authorizer, err := auth.NewAuthorizerFromFile(fs.Values["ResourceManagerEndpoint"])
	if err != nil {
		fmt.Println("Failed to get OAuth config:", err)
	}
	subsID := fs.Values["AZURE_SUBSCRIPTION_ID"]
	tenantID := fs.Values["AZURE_TENANT_ID"]

	fmt.Println(subsID)
	fmt.Println(tenantID)


	sess := AzureSession {
		SubscriptionId: subsID,
		TenantId:		tenantID,
		Authorizer:     authorizer,
	}

	return &sess, nil
}
