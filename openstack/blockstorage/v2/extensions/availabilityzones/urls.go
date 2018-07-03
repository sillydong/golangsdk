package availabilityzones

import "github.com/huaweicloud/golangsdk"

func listURL(c *golangsdk.ServiceClient) string {
	return c.ServiceURL("os-availability-zone")
}
