package tenants

import "github.com/huaweicloud/golangsdk"

// listURL generates URL for listing tenants.
func listURL(client *golangsdk.ServiceClient) string {
	return client.ServiceURL("tenants")
}

// getURL generates URL for getting tenant.
func getURL(client *golangsdk.ServiceClient, tenantID string) string {
	return client.ServiceURL("tenants", tenantID)
}

// createURL generates URL for createing tenant.
func createURL(client *golangsdk.ServiceClient) string {
	return client.ServiceURL("tenants")
}

// deleteURL generates URL for deleteing tenant.
func deleteURL(client *golangsdk.ServiceClient, tenantID string) string {
	return client.ServiceURL("tenants", tenantID)
}

// updateURL generates URL for updating a tenant.
func updateURL(client *golangsdk.ServiceClient, tenantID string) string {
	return client.ServiceURL("tenants", tenantID)
}
