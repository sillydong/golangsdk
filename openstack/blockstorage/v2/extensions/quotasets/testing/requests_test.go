package testing

import (
	"testing"

	"github.com/huaweicloud/golangsdk/openstack/blockstorage/v2/extensions/quotasets"
	th "github.com/huaweicloud/golangsdk/testhelper"
	"github.com/huaweicloud/golangsdk/testhelper/client"
)

func TestGetUsage(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	MockGetUsageResponse(t)

	v, err := quotasets.GetUsage(client.ServiceClient(), FirstTenantID).Extract()
	th.AssertNoErr(t, err)

	th.AssertEquals(t, v.ID, "cd631140887d4b6e9c786b67a6dd4c02")
}
