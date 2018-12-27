package floatingips

import "github.com/huaweicloud/golangsdk"

const resourcePath = "os-floating-ips"

func resourceURL(c *golangsdk.ServiceClient) string {
	return c.ServiceURL(resourcePath)
}

func listURL(c *golangsdk.ServiceClient) string {
	return resourceURL(c)
}

func createURL(c *golangsdk.ServiceClient) string {
	return resourceURL(c)
}

func getURL(c *golangsdk.ServiceClient, id string) string {
	return c.ServiceURL(resourcePath, id)
}

func deleteURL(c *golangsdk.ServiceClient, id string) string {
	return getURL(c, id)
}

func serverURL(c *golangsdk.ServiceClient, serverID string) string {
	return c.ServiceURL("servers/" + serverID + "/action")
}

// associateURL generate a url to associate floating ip with server
func associateURL(c *golangsdk.ServiceClient, serverID string) string {
	return serverURL(c, serverID)
}

// disassociateURL generate a url to disassociate floating ip with server
func disassociateURL(c *golangsdk.ServiceClient, serverID string) string {
	return serverURL(c, serverID)
}
