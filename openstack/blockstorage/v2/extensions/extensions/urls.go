package extensions

import "github.com/huaweicloud/golangsdk"

// ListExtensionURL generates the URL for the extensions resource collection.
func ListExtensionURL(c *golangsdk.ServiceClient) string {
	return c.ServiceURL("extensions")
}
