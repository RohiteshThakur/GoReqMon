package azure

import (
	"fmt"
	Azure "github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/azure/auth"
	//"bytes"
	//"github.com/marstr/randname"
)

var (
	// these are our *global* config settings, to be shared by all packages.
	// each has corresponding public accessors below.
	// if anything requires a `Set` accessor, that indicates it perhaps
	// shouldn't be set here, because mutable vars shouldn't be global.
	clientID                string 								// Set in AUTH File.
	clientSecret            string								// Set in AUTH File.
	tenantID                string 								// Set in AUTH File.
	subscriptionID          string 								// Set in AUTH File.
	groupName               string 			    				// deprecated, use baseGroupName instead
	locationDefault        	string = "West Europe"
	authorizationServerURL 	string = "https://login.microsoftonline.com"
	cloudName              	string = "AzurePublicCloud"
	useDeviceFlow          	bool   = false
	keepResources          	bool   = false
	baseGroupName          	string = "<base group name>"		// update:
	userAgent              	string								// see below
	environment             *Azure.Environment 					// var environments = map[string]Environment { "AZUREPUBLICCLOUD":       PublicCloud, }
)


// DefaultLocation() returns the default location wherein to create new resources.
// Some resource types are not available in all locations so another location might need
// to be chosen.

func GetClientID (fs auth.FileSettings) (string) {
	clientID = fs.Values["AZURE_CLIENT_ID"]
	return clientID
}

func GetClientSecret (fs auth.FileSettings) (string) {
	clientSecret = fs.Values["AZURE_CLIENT_SECRET"]
	return clientSecret
}

func GetSubscriptionId (fs auth.FileSettings) (string) {
	subscriptionID = fs.Values["AZURE_SUBSCRIPTION_ID"]
	return subscriptionID
}

func GetTenantId (fs auth.FileSettings) (string) {
	tenantID = fs.Values["AZURE_TENANT_ID"]
	return tenantID
}

func DefaultLocation() string {
	return locationDefault
}

// AuthorizationServerURL is the OAuth authorization server URL.
// Q: Can this be gotten from the `azure.Environment` in `Environment()`?
func AuthorizationServerURL() string {
	return authorizationServerURL
}

// UseDeviceFlow() specifies if interactive auth should be used. Interactive
// auth uses the OAuth Device Flow grant type.
func UseDeviceFlow() bool {
	return useDeviceFlow
}

// deprecated: do not use global group names
// utilize `BaseGroupName()` for a shared prefix
func GroupName() string {
	return groupName
}

// deprecated: we have to set this because we use a global for group names
// once that's fixed this should be removed
func SetGroupName(name string) {
	groupName = name
}

// BaseGroupName() returns a prefix for new groups.
func BaseGroupName() string {
	return baseGroupName
}

// KeepResources() specifies whether to keep resources created by samples.
func KeepResources() bool {
	return keepResources
}

// UserAgent() specifies a string to append to the agent identifier.
func UserAgent() string {
	if len(userAgent) > 0 {
		return userAgent
	}
	return "<base-group-name>"													// update:
}

// Environment() returns an `azure.Environment{...}` for the current cloud.
func Environment() *Azure.Environment {
	if environment != nil {
		return environment
	}
	env, err := Azure.EnvironmentFromName(cloudName)
	if err != nil {
		// TODO: move to initialization of var
		panic(fmt.Sprintf(
			"invalid cloud name '%s' specified, cannot continue\n", cloudName))
	}
	environment = &env
	return environment
}

/*
// GenerateGroupName leverages BaseGroupName() to return a more detailed name,
// helping to avoid collisions.  It appends each of the `affixes` to
// BaseGroupName() separated by dashes, and adds a 5-character random string.
func GenerateGroupName(affixes ...string) string {
	// go1.10+
	// import strings
	// var b strings.Builder
	// b.WriteString(BaseGroupName())
	b := bytes.NewBufferString(BaseGroupName())
	b.WriteRune('-')
	for _, affix := range affixes {
		b.WriteString(affix)
		b.WriteRune('-')
	}
	return randname.GenerateWithPrefix(b.String(), 5)
}

// AppendRandomSuffix will append a suffix of five random characters to the specified prefix.
func AppendRandomSuffix(prefix string) string {
	return randname.GenerateWithPrefix(prefix, 5)
}
*/