package availabilityzones

import "github.com/huaweicloud/golangsdk"

// listURL generate url to list availability zones
func listURL(c *golangsdk.ServiceClient) string {
	return c.ServiceURL("os-availability-zone")
}

// func listDetailURL(c *golangsdk.ServiceClient) string {
//	return c.ServiceURL("os-availability-zone", "detail")
// }
