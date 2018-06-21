// +build acceptance imageservice images

package v2

import (
	"testing"

	"github.com/huaweicloud/golangsdk/acceptance/clients"
	"github.com/huaweicloud/golangsdk/acceptance/tools"
	"github.com/huaweicloud/golangsdk/openstack/imageservice/v1/cloudimages"
	"github.com/huaweicloud/golangsdk/openstack/imageservice/v2/images"
	"github.com/huaweicloud/golangsdk/pagination"
	th "github.com/huaweicloud/golangsdk/testhelper"
)

func TestOperateImageTags(t *testing.T) {

	clientv2, err := clients.NewImageServiceV2Client()
	th.AssertNoErr(t, err)

	imageID := ""
	err = images.List(clientv2, images.ListOpts{
		Limit: 1,
	}).EachPage(func(page pagination.Page) (bool, error) {
		images, err := images.ExtractImages(page)
		if err != nil {
			t.Fatalf("Unable to extract images: %v", err)
		}

		for _, image := range images {
			imageID = image.ID
			return false, nil
		}

		return true, nil
	})
	tools.PrintResource(t, imageID)

	client, err := clients.NewImageServiceV1Client()
	th.AssertNoErr(t, err)

	// FIXME broken, caused by huaweicloud platform
	setTagResult := cloudimages.SetImageTag(client, cloudimages.SetImageTagOpts{
		ImageID: imageID,
		Tag:     "golangtag.blabla",
	})
	th.AssertNoErr(t, setTagResult.Err)

	listTagsResult := cloudimages.ListImageTags(client, cloudimages.ListImageTagsOpts{
		Limit: 10,
		Page:  1,
	})

	tags, err := listTagsResult.Extract()
	th.AssertNoErr(t, err)

	tools.PrintResource(t, tags)
}
