package snapshots

import "github.com/huaweicloud/golangsdk"

// 5.1
func createURL(c *golangsdk.ServiceClient) string {
	return c.ServiceURL("snapshots")
}

// 5.3
func deleteURL(c *golangsdk.ServiceClient, id string) string {
	return c.ServiceURL("snapshots", id)
}

// 5.11
func getURL(c *golangsdk.ServiceClient, id string) string {
	return deleteURL(c, id)
}

// 5.5
func updateURL(c *golangsdk.ServiceClient, id string) string {
	return deleteURL(c, id)
}

// 5.7
func listURL(c *golangsdk.ServiceClient) string {
	return createURL(c)
}

// 5.9
func detailURL(c *golangsdk.ServiceClient) string {
	return c.ServiceURL("snapshots", "detail")
}

// 5.15 5.17 5.21
func metadataURL(c *golangsdk.ServiceClient, id string) string {
	return c.ServiceURL("snapshots", id, "metadata")
}

// 5.19 5.25 5.23
func metadataKeyURL(c *golangsdk.ServiceClient, id, key string) string {
	return c.ServiceURL("snapshots", id, "metadata", key)
}

// 5.13
func rollbackURL(c *golangsdk.ServiceClient, id string) string {
	return c.ServiceURL("os-vendor-snapshots", id, "rollback")
}
