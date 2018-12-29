package projects

import "github.com/huaweicloud/golangsdk"

// listURL generate url to querying information about a specified project
func listURL(client *golangsdk.ServiceClient) string {
	return client.ServiceURL("projects")
}

// getURL generate url to show project details
func getURL(client *golangsdk.ServiceClient, projectID string) string {
	return client.ServiceURL("projects", projectID)
}

// createURL generate url to add project
func createURL(client *golangsdk.ServiceClient) string {
	return client.ServiceURL("projects")
}

func deleteURL(client *golangsdk.ServiceClient, projectID string) string {
	return client.ServiceURL("projects", projectID)
}

// updateURL generate url to update project
func updateURL(client *golangsdk.ServiceClient, projectID string) string {
	return client.ServiceURL("projects", projectID)
}
