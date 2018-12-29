package roles

import "github.com/huaweicloud/golangsdk"

const (
	rolePath = "roles"
)

// listURL generate url to list roles
func listURL(client *golangsdk.ServiceClient) string {
	return client.ServiceURL(rolePath)
}

// getURL generate url to show role details
func getURL(client *golangsdk.ServiceClient, roleID string) string {
	return client.ServiceURL(rolePath, roleID)
}

func createURL(client *golangsdk.ServiceClient) string {
	return client.ServiceURL(rolePath)
}

func updateURL(client *golangsdk.ServiceClient, roleID string) string {
	return client.ServiceURL(rolePath, roleID)
}

func deleteURL(client *golangsdk.ServiceClient, roleID string) string {
	return client.ServiceURL(rolePath, roleID)
}

func listAssignmentsURL(client *golangsdk.ServiceClient) string {
	return client.ServiceURL("role_assignments")
}

// assignURL generate url to assign or unassign role to group on domain or
// project
func assignURL(client *golangsdk.ServiceClient, targetType, targetID, actorType, actorID, roleID string) string {
	return client.ServiceURL(targetType, targetID, actorType, actorID, rolePath, roleID)
}

// checkRoleOfURL generate url to check whether group has role assignment on
// domain or project
func checkRoleOfURL(client *golangsdk.ServiceClient,
	targetType, targetID, actorType, actorID, roleID string) string {

	return client.ServiceURL(
		targetType, targetID,
		actorType, actorID,
		"roles", roleID,
	)
}

// listRolesOfURL generate url to list rolesto group on domain or project
func listRolesOfURL(client *golangsdk.ServiceClient,
	targetType, targetID, actorType, actorID string) string {

	return client.ServiceURL(
		targetType, targetID,
		actorType, actorID,
		"roles",
	)
}
