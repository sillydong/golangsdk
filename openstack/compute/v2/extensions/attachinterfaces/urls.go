package attachinterfaces

import "github.com/huaweicloud/golangsdk"

// listInterfaceURL generate URL to list port interfaces
func listInterfaceURL(client *golangsdk.ServiceClient, serverID string) string {
	return client.ServiceURL("servers", serverID, "os-interface")
}

// getInterfaceURL generate URL to show port interface details
func getInterfaceURL(client *golangsdk.ServiceClient, serverID, portID string) string {
	return client.ServiceURL("servers", serverID, "os-interface", portID)
}

// createInterfaceURL generate URL to attach interface
func createInterfaceURL(client *golangsdk.ServiceClient, serverID string) string {
	return client.ServiceURL("servers", serverID, "os-interface")
}

// deleteInterfaceURL generate URL to detach interface
func deleteInterfaceURL(client *golangsdk.ServiceClient, serverID, portID string) string {
	return client.ServiceURL("servers", serverID, "os-interface", portID)
}
