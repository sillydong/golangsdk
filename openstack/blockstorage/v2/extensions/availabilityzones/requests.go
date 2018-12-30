package availabilityzones

import (
	"github.com/huaweicloud/golangsdk"
	"github.com/huaweicloud/golangsdk/pagination"
)

// List will return the existing availability zones.
func List(client *golangsdk.ServiceClient) pagination.Pager {
	return pagination.NewPager(
		client,
		listURL(client),
		func(r pagination.PageResult) pagination.Page {
			return AvailabilityZonePage{pagination.SinglePageBase(r)}
		},
	)
}
