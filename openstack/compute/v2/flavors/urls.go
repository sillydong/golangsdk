package flavors

import "github.com/huaweicloud/golangsdk"

// getURL generate URL to get flavor detail
func getURL(client *golangsdk.ServiceClient, id string) string {
	return client.ServiceURL("flavors", id)
}

// detailURL generate URL for listing flavor details
func detailURL(client *golangsdk.ServiceClient) string {
	return client.ServiceURL("flavors", "detail")
}

// listURL generate URL to list flavors
func listURL(client *golangsdk.ServiceClient) string {
	return client.ServiceURL("flavors")
}

// func createURL(client *golangsdk.ServiceClient) string {
// 	return client.ServiceURL("flavors")
// }

// func deleteURL(client *golangsdk.ServiceClient, id string) string {
// 	return client.ServiceURL("flavors", id)
// }

// func accessURL(client *golangsdk.ServiceClient, id string) string {
// 	return client.ServiceURL("flavors", id, "os-flavor-access")
// }

// func accessActionURL(client *golangsdk.ServiceClient, id string) string {
// 	return client.ServiceURL("flavors", id, "action")
// }

// extraSpecsListURL generate URL to list all extra specs for a flavor
func extraSpecsListURL(client *golangsdk.ServiceClient, id string) string {
	return client.ServiceURL("flavors", id, "os-extra_specs")
}

// func extraSpecsGetURL(client *golangsdk.ServiceClient, id, key string) string {
// 	return client.ServiceURL("flavors", id, "os-extra_specs", key)
// }

// func extraSpecsCreateURL(client *golangsdk.ServiceClient, id string) string {
// 	return client.ServiceURL("flavors", id, "os-extra_specs")
// }

// func extraSpecUpdateURL(client *golangsdk.ServiceClient, id, key string) string {
// 	return client.ServiceURL("flavors", id, "os-extra_specs", key)
// }

// func extraSpecDeleteURL(client *golangsdk.ServiceClient, id, key string) string {
// 	return client.ServiceURL("flavors", id, "os-extra_specs", key)
// }
