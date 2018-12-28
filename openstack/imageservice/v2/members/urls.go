package members

import "github.com/huaweicloud/golangsdk"

func imageMembersURL(c *golangsdk.ServiceClient, imageID string) string {
	return c.ServiceURL("images", imageID, "members")
}

// listMembersURL generate a URL for listing all members of an image
func listMembersURL(c *golangsdk.ServiceClient, imageID string) string {
	return imageMembersURL(c, imageID)
}

// createMemberURL generate a URL for adding a member to an image
func createMemberURL(c *golangsdk.ServiceClient, imageID string) string {
	return imageMembersURL(c, imageID)
}

// imageMemberURL generate a URL for  getting a image member
func imageMemberURL(c *golangsdk.ServiceClient, imageID string, memberID string) string {
	return c.ServiceURL("images", imageID, "members", memberID)
}

// getMemberURL generate a URL to get the detail of an image member
func getMemberURL(c *golangsdk.ServiceClient, imageID string, memberID string) string {
	return imageMemberURL(c, imageID, memberID)
}

// updateMemberURL generate a URL to update the status of an image member
func updateMemberURL(c *golangsdk.ServiceClient, imageID string, memberID string) string {
	return imageMemberURL(c, imageID, memberID)
}

// deleteMemberURL generate a URL to remove an image member
func deleteMemberURL(c *golangsdk.ServiceClient, imageID string, memberID string) string {
	return imageMemberURL(c, imageID, memberID)
}

// getMemberSchemas generate a url to get member schemas
func getMemberSchemas(c *golangsdk.ServiceClient) string {
	return c.ServiceURL("schemas", "member")
}

// getMembersSchemas generate a url to get member schemas
func getMembersSchemas(c *golangsdk.ServiceClient) string {
	return c.ServiceURL("schemas", "members")
}
