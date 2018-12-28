package availabilityzones

import "github.com/huaweicloud/golangsdk"

// listURL generates URL for list avaliabilityzones
func listURL(c *golangsdk.ServiceClient) string {
	return c.ServiceURL("os-availability-zone")
}
