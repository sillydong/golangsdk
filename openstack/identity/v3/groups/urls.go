package groups

import "github.com/huaweicloud/golangsdk"

// listURL generate url to list groups
func listURL(client *golangsdk.ServiceClient) string {
	return client.ServiceURL("groups")
}

// getURL generate url to show group details
func getURL(client *golangsdk.ServiceClient, groupID string) string {
	return client.ServiceURL("groups", groupID)
}

// createURL generate url to create group
func createURL(client *golangsdk.ServiceClient) string {
	return client.ServiceURL("groups")
}

// updateURl generate url to update group
func updateURL(client *golangsdk.ServiceClient, groupID string) string {
	return client.ServiceURL("groups", groupID)
}

// deleteURL generate url to delete group
func deleteURL(client *golangsdk.ServiceClient, groupID string) string {
	return client.ServiceURL("groups", groupID)
}
