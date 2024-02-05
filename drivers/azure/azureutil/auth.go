package azureutil

import (
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Azure driver allows two authentication methods:
//
// 1. OAuth Device Flow
//
// Azure Active Directory implements OAuth 2.0 Device Flow described here:
// https://tools.ietf.org/html/draft-denniss-oauth-device-flow-00. It is simple
// for users to authenticate through a browser and requires re-authenticating
// every 2 weeks.
//
// Device auth prints a message to the screen telling the user to click on URL
// and approve the app on the browser, meanwhile the client polls the auth API
// for a token. Once we have token, we save it locally to a file with proper
// permissions and when the token expires (in Azure case typically 1 hour) SDK
// will automatically refresh the specified token and will call the refresh
// callback function we implement here. This way we will always be storing a
// token with a refresh_token saved on the machine.
//
// 2. Azure Service Principal Account
//
// This is designed for headless authentication to Azure APIs but requires more
// steps from user to create a Service Principal Account and provide its
// credentials to the machine driver.

var (
	// AD app id for docker-machine driver in various Azure realms
	appIDs = map[string]string{
		azure.PublicCloud.Name: "637ddaba-219b-43b8-bf19-8cea500cf273",
		azure.ChinaCloud.Name:  "bb5eed6f-120b-4365-8fd9-ab1a3fba5698",
		azure.GermanCloud.Name: "aabac5f7-dd47-47ef-824c-e0d57598cada",
	}
)

// AuthenticateDeviceFlow fetches a token from the local file cache or initiates a consent
// flow and waits for token to be obtained. Obtained token is stored in a file cache for
// future use and refreshing.
func AuthenticateDeviceFlow() (azcore.TokenCredential, error) {
	cred, err := azidentity.NewDeviceCodeCredential(&azidentity.DeviceCodeCredentialOptions{})
	return cred, err
}

// AuthenticateServicePrincipal uses given service principal credentials to return a
// service principal token. Generated token is not stored in a cache file or refreshed.
func AuthenticateServicePrincipal(tenantID, clientID, clientSecret string) (azcore.TokenCredential, error) {
	cred, err := azidentity.NewClientSecretCredential(tenantID, clientID, clientSecret, &azidentity.ClientSecretCredentialOptions{})
	return cred, err
}
