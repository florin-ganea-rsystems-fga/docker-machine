package azureutil

import (
	"fmt"
	"github.com/docker/machine/version"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	compute "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute"
	network "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork"
	resources "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armresources"
	subscriptions "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armsubscriptions"
	storage "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storage/armstorage"
	"github.com/Azure/go-autorest/autorest"
)

// TODO(ahmetalpbalkan) Remove duplication around client creation. This is
// happening because we auto-generate our SDK and we don't have generics in Go.
// We are hoping to come up with a factory or some defaults instance to set
// these client configuration in a central place in azure-sdk-for-go.

func oauthClient() autorest.Client {
	c := autorest.NewClientWithUserAgent(fmt.Sprintf("docker-machine/%s", version.Version))
	c.RequestInspector = withInspection()
	c.ResponseInspector = byInspecting()
	return c
}

func subscriptionsClient(credential azcore.TokenCredential) *subscriptions.SubscriptionClient {
	c, _ := subscriptions.NewSubscriptionClient(credential, &arm.ClientOptions{})
	return c
}

func (a AzureClient) providersClient() *resources.ProvidersClient {
	c, _ := resources.NewProvidersClient(a.subscriptionID, a.auth, &arm.ClientOptions{})
	return c
}

func (a AzureClient) resourceGroupsClient() *resources.ResourceGroupsClient {
	c, _ := resources.NewResourceGroupsClient(a.subscriptionID, a.auth, &arm.ClientOptions{})
	return c
}

func (a AzureClient) securityGroupsClient() *network.SecurityGroupsClient {
	c, _ := network.NewSecurityGroupsClient(a.subscriptionID, a.auth, &arm.ClientOptions{})
	return c
}

func (a AzureClient) virtualNetworksClient() *network.VirtualNetworksClient {
	c, _ := network.NewVirtualNetworksClient(a.subscriptionID, a.auth, &arm.ClientOptions{})
	return c
}

func (a AzureClient) subnetsClient() *network.SubnetsClient {
	c, _ := network.NewSubnetsClient(a.subscriptionID, a.auth, &arm.ClientOptions{})
	return c
}

func (a AzureClient) networkInterfacesClient() *network.InterfacesClient {
	c, _ := network.NewInterfacesClient(a.subscriptionID, a.auth, &arm.ClientOptions{})
	return c
}

func (a AzureClient) publicIPAddressClient() *network.PublicIPAddressesClient {
	c, _ := network.NewPublicIPAddressesClient(a.subscriptionID, a.auth, &arm.ClientOptions{})
	return c
}

func (a AzureClient) storageAccountsClient() *storage.AccountsClient {
	c, _ := storage.NewAccountsClient(a.subscriptionID, a.auth, &arm.ClientOptions{})
	return c
}

func (a AzureClient) virtualMachinesClient() *compute.VirtualMachinesClient {
	c, _ := compute.NewVirtualMachinesClient(a.subscriptionID, a.auth, &arm.ClientOptions{})
	return c
}

func (a AzureClient) availabilitySetsClient() *compute.AvailabilitySetsClient {
	c, _ := compute.NewAvailabilitySetsClient(a.subscriptionID, a.auth, &arm.ClientOptions{})
	return c
}
