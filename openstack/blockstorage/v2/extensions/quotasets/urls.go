package quotasets

import "github.com/huaweicloud/golangsdk"

const resourcePath = "os-quota-sets"

func getURL(c *golangsdk.ServiceClient, projectID string) string {
	return c.ServiceURL(resourcePath, projectID)
}
