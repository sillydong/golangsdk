package volumetypes

import (
	"github.com/huaweicloud/golangsdk"
)

// listURL generate URL for list volumetypes.
func listURL(c *golangsdk.ServiceClient) string {
	return c.ServiceURL("types")
}

// getURL generate URL for get volumetype with id.
func getURL(c *golangsdk.ServiceClient, id string) string {
	return c.ServiceURL("types", id)
}
