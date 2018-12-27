package servers

import "github.com/huaweicloud/golangsdk"

// createURL generate url to create server
func createURL(client *golangsdk.ServiceClient) string {
	return client.ServiceURL("servers")
}

// listURL generate url for querying ECSes
func listURL(client *golangsdk.ServiceClient) string {
	return createURL(client)
}

// listDetailURL generate url to list servers details
func listDetailURL(client *golangsdk.ServiceClient) string {
	return client.ServiceURL("servers", "detail")
}

// deleteURL generate url to delete server
func deleteURL(client *golangsdk.ServiceClient, id string) string {
	return client.ServiceURL("servers", id)
}

// getURL generate url to get one server detail
func getURL(client *golangsdk.ServiceClient, id string) string {
	return deleteURL(client, id)
}

// func updateURL(client *golangsdk.ServiceClient, id string) string {
// 	return deleteURL(client, id)
// }

// actionURL generate action URL for server starting, rebooting and stopping
func actionURL(client *golangsdk.ServiceClient, id string) string {
	return client.ServiceURL("servers", id, "action")
}

// metadatumURL generate URL to get or set metadata items
func metadatumURL(client *golangsdk.ServiceClient, id, key string) string {
	return client.ServiceURL("servers", id, "metadata", key)
}

// metadataURL generate URL to list or update metadatas
func metadataURL(client *golangsdk.ServiceClient, id string) string {
	return client.ServiceURL("servers", id, "metadata")
}

// func listAddressesURL(client *golangsdk.ServiceClient, id string) string {
// 	return client.ServiceURL("servers", id, "ips")
// }

// func listAddressesByNetworkURL(client *golangsdk.ServiceClient, id, network string) string {
// 	return client.ServiceURL("servers", id, "ips", network)
// }

// func passwordURL(client *golangsdk.ServiceClient, id string) string {
// 	return client.ServiceURL("servers", id, "os-server-password")
// }
