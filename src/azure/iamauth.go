package azure

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/adal"
	"github.com/Azure/go-autorest/autorest/azure/auth"
)

var (
	armAuthorizer      autorest.Authorizer
	batchAuthorizer    autorest.Authorizer
	graphAuthorizer    autorest.Authorizer
	keyvaultAuthorizer autorest.Authorizer
)

// OAuthGrantType specifies which grant type to use.
type OAuthGrantType int

const (
	// OAuthGrantTypeServicePrincipal for client credentials flow
	OAuthGrantTypeServicePrincipal OAuthGrantType = iota
	// OAuthGrantTypeDeviceFlow for device flow
	OAuthGrantTypeDeviceFlow
)

// GrantType returns what grant type has been configured.
func grantType() OAuthGrantType {
	if UseDeviceFlow() {
		return OAuthGrantTypeDeviceFlow
	}
	return OAuthGrantTypeServicePrincipal
}

// GetResourceManagementAuthorizer gets an OAuthTokenAuthorizer for Azure Resource Manager
func GetResourceManagementAuthorizer(fs auth.FileSettings) (autorest.Authorizer, error) {
	if armAuthorizer != nil {
		return armAuthorizer, nil
	}

	var a autorest.Authorizer
	var err error

	a, err = getAuthorizerForResource(
		grantType(), Environment().ResourceManagerEndpoint, fs)

	if err == nil {
		// cache
		armAuthorizer = a
	} else {
		// clear cache
		armAuthorizer = nil
	}
	return armAuthorizer, err
}

// GetBatchAuthorizer gets an OAuthTokenAuthorizer for Azure Batch.
func GetBatchAuthorizer(fs auth.FileSettings) (autorest.Authorizer, error) {
	if batchAuthorizer != nil {
		return batchAuthorizer, nil
	}

	var a autorest.Authorizer
	var err error

	a, err = getAuthorizerForResource(
		grantType(), Environment().BatchManagementEndpoint, fs)

	if err == nil {
		// cache
		batchAuthorizer = a
	} else {
		// clear cache
		batchAuthorizer = nil
	}

	return batchAuthorizer, err
}

// GetGraphAuthorizer gets an OAuthTokenAuthorizer for graphrbac API.
func GetGraphAuthorizer(fs auth.FileSettings) (autorest.Authorizer, error) {
	if graphAuthorizer != nil {
		return graphAuthorizer, nil
	}

	var a autorest.Authorizer
	var err error

	a, err = getAuthorizerForResource(grantType(), Environment().GraphEndpoint, fs)

	if err == nil {
		// cache
		graphAuthorizer = a
		fmt.Println("Authorization successful.")
	} else {
		graphAuthorizer = nil
		fmt.Println ("Authorization failed.")
	}

	return graphAuthorizer, err
}

// GetKeyvaultAuthorizer gets an OAuthTokenAuthorizer for use with Key Vault
// keys and secrets. Note that Key Vault *Vaults* are managed by Azure Resource
// Manager.
func GetKeyvaultAuthorizer(fs auth.FileSettings) (autorest.Authorizer, error) {
	if keyvaultAuthorizer != nil {
		return keyvaultAuthorizer, nil
	}

	// BUG: default value for KeyVaultEndpoint is wrong
	vaultEndpoint := strings.TrimSuffix(Environment().KeyVaultEndpoint, "/")
	// BUG: alternateEndpoint replaces other endpoints in the configs below
	alternateEndpoint, _ := url.Parse(
		"https://login.windows.net/" + GetTenantId(fs) + "/oauth2/token")

	var a autorest.Authorizer
	var err error

	switch grantType() {
	case OAuthGrantTypeServicePrincipal:
		oauthconfig, err := adal.NewOAuthConfig(
			Environment().ActiveDirectoryEndpoint, GetTenantId(fs))
		if err != nil {
			return a, err
		}
		oauthconfig.AuthorizeEndpoint = *alternateEndpoint

		token, err := adal.NewServicePrincipalToken(
			*oauthconfig, GetClientID(fs), GetClientSecret(fs), vaultEndpoint)
		if err != nil {
			return a, err
		}

		a = autorest.NewBearerAuthorizer(token)

	case OAuthGrantTypeDeviceFlow:
		deviceConfig := auth.NewDeviceFlowConfig(GetClientID(fs), GetTenantId(fs))
		deviceConfig.Resource = vaultEndpoint
		deviceConfig.AADEndpoint = alternateEndpoint.String()
		a, err = deviceConfig.Authorizer()
	default:
		return a, fmt.Errorf("invalid grant type specified")
	}

	if err == nil {
		keyvaultAuthorizer = a
	} else {
		keyvaultAuthorizer = nil
	}

	return keyvaultAuthorizer, err
}

func getAuthorizerForResource(grantType OAuthGrantType, resource string, fs auth.FileSettings) (autorest.Authorizer, error) {

	var a autorest.Authorizer
	var err error

	switch grantType {

	case OAuthGrantTypeServicePrincipal:
		oauthConfig, err := adal.NewOAuthConfig(
			Environment().ActiveDirectoryEndpoint, GetTenantId(fs))
		if err != nil {
			return nil, err
		}

		token, err := adal.NewServicePrincipalToken(
			*oauthConfig, GetClientID(fs), GetClientSecret(fs), resource)
		if err != nil {
			return nil, err
		}
		a = autorest.NewBearerAuthorizer(token)

	case OAuthGrantTypeDeviceFlow:
		deviceconfig := auth.NewDeviceFlowConfig(GetClientID(fs), GetTenantId(fs))
		deviceconfig.Resource = resource
		a, err = deviceconfig.Authorizer()
		if err != nil {
			return nil, err
		}

	default:
		return a, fmt.Errorf("invalid grant type specified")
	}

	return a, err
}

// GetResourceManagementTokenHybrid retrieves auth token for hybrid environment
func GetResourceManagementTokenHybrid(activeDirectoryEndpoint, tokenAudience string, fs auth.FileSettings) (adal.OAuthTokenProvider, error) {
	var tokenProvider adal.OAuthTokenProvider
	oauthConfig, err := adal.NewOAuthConfig(activeDirectoryEndpoint, GetTenantId(fs))
	tokenProvider, err = adal.NewServicePrincipalToken(
		*oauthConfig,
		GetClientID(fs),
		GetClientSecret(fs),
		tokenAudience)

	return tokenProvider, err
}