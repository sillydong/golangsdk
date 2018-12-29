package endpoints

import "github.com/huaweicloud/golangsdk"

// listURL generate url to querying endpoints
func listURL(client *golangsdk.ServiceClient) string {
	return client.ServiceURL("endpoints")
}

// endpointURL generate url to show endpoint details
func endpointURL(client *golangsdk.ServiceClient, endpointID string) string {
	return client.ServiceURL("endpoints", endpointID)
}
