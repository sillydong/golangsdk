package volumeattach

import "github.com/huaweicloud/golangsdk"

const resourcePath = "os-volume_attachments"

func resourceURL(c *golangsdk.ServiceClient, serverID string) string {
	return c.ServiceURL("servers", serverID, resourcePath)
}

// listURL generate URL to list attached volumes
func listURL(c *golangsdk.ServiceClient, serverID string) string {
	return resourceURL(c, serverID)
}

// createURL generate URL to attach volume
func createURL(c *golangsdk.ServiceClient, serverID string) string {
	return resourceURL(c, serverID)
}

// getURL generate URL to show attached volumes
func getURL(c *golangsdk.ServiceClient, serverID, aID string) string {
	return c.ServiceURL("servers", serverID, resourcePath, aID)
}

// deleteURL generate URL to detach volume
func deleteURL(c *golangsdk.ServiceClient, serverID, aID string) string {
	return getURL(c, serverID, aID)
}
