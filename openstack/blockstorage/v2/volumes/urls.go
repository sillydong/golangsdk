package volumes

import "github.com/huaweicloud/golangsdk"

// 3.3
func createURL(c *golangsdk.ServiceClient) string {
	return c.ServiceURL("volumes")
}

// 3.13
func listURL(c *golangsdk.ServiceClient) string {
	return c.ServiceURL("volumes")
}

// 3.18
func detailURL(c *golangsdk.ServiceClient) string {
	return c.ServiceURL("volumes", "detail")
}

// 3.7
func deleteURL(c *golangsdk.ServiceClient, id string) string {
	return c.ServiceURL("volumes", id)
}

// 3.23
func getURL(c *golangsdk.ServiceClient, id string) string {
	return deleteURL(c, id)
}

// 3.10
func updateURL(c *golangsdk.ServiceClient, id string) string {
	return deleteURL(c, id)
}
