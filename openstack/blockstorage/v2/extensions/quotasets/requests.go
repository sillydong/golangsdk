package quotasets

import (
	"fmt"

	"github.com/huaweicloud/golangsdk"
)

// GetUsage returns detailed public data about a previously created QuotaSet.
func GetUsage(client *golangsdk.ServiceClient, projectID string) (r GetUsageResult) {
	u := fmt.Sprintf("%s?usage=true", getURL(client, projectID))
	_, r.Err = client.Get(u, &r.Body, nil)
	return
}
