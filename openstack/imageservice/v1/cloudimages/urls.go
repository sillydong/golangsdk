package cloudimages

import (
	"github.com/huaweicloud/golangsdk"
)

func getImageTagsURL(c *golangsdk.ServiceClient) string {
	return c.ServiceURL("cloudimages", "tags")
}

func putImageTagsURL(c *golangsdk.ServiceClient) string {
	return c.ServiceURL("cloudimages", "tags")
}

func importImageURL(c *golangsdk.ServiceClient, imageID string) string {
	return c.ServiceURL("cloudimages", imageID, "upload")
}

func exportImageURL(c *golangsdk.ServiceClient, imageID string) string {
	return c.ServiceURL("cloudimages", imageID, "file")
}

func copyImageURL(c *golangsdk.ServiceClient, imageID string) string {
	return c.ServiceURL("cloudimages", imageID, "copy")
}

func imageMemberOpURL(c *golangsdk.ServiceClient) string {
	return c.ServiceURL("cloudimages", "members")
}
