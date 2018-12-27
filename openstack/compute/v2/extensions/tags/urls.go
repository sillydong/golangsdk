package tags

import "github.com/huaweicloud/golangsdk"

func operateTagList(client *golangsdk.ServiceClient, id string) string {
	return client.ServiceURL("servers", id, "tags")
}

var (
	// listTags generate url to list tags
	listTags = operateTagList
	// replaceTags generate url to replace tags
	replaceTags = operateTagList
	// deleteTags generate url to delete all tags
	deleteTags = operateTagList
)

func operateTag(client *golangsdk.ServiceClient, id, tag string) string {
	return client.ServiceURL("servers", id, "tags", tag)
}

var (
	// checkTag generate url to check tag existence
	checkTag = operateTag
	// addTag generate url to add a single tag
	addTag = operateTag
	// deleteTag generate url to delete a single tag
	deleteTag = operateTag
)
