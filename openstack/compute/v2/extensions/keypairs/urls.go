package keypairs

import "github.com/huaweicloud/golangsdk"

const resourcePath = "os-keypairs"

func resourceURL(c *golangsdk.ServiceClient) string {
	return c.ServiceURL(resourcePath)
}

// listURL generate url to list keypairs
func listURL(c *golangsdk.ServiceClient) string {
	return resourceURL(c)
}

// createURL generate url to add keypair
func createURL(c *golangsdk.ServiceClient) string {
	return resourceURL(c)
}

// getURl generate url to get keypairs
func getURL(c *golangsdk.ServiceClient, name string) string {
	return c.ServiceURL(resourcePath, name)
}

// deleteURL generate url to delete keypair
func deleteURL(c *golangsdk.ServiceClient, name string) string {
	return getURL(c, name)
}
