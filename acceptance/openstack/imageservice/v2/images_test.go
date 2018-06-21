// +build acceptance imageservice images

package v2

import (
	"testing"
	"time"

	"github.com/huaweicloud/golangsdk/acceptance/clients"
	"github.com/huaweicloud/golangsdk/acceptance/tools"
	"github.com/huaweicloud/golangsdk/openstack/imageservice/v2/images"
	"github.com/huaweicloud/golangsdk/pagination"
	th "github.com/huaweicloud/golangsdk/testhelper"
)

func TestImagesListEachPage(t *testing.T) {
	client, err := clients.NewImageServiceV2Client()
	th.AssertNoErr(t, err)

	listOpts := images.ListOpts{
		Limit: 1,
	}

	pager := images.List(client, listOpts)
	err = pager.EachPage(func(page pagination.Page) (bool, error) {
		images, err := images.ExtractImages(page)
		if err != nil {
			t.Fatalf("Unable to extract images: %v", err)
		}

		for _, image := range images {
			tools.PrintResource(t, image)
			tools.PrintResource(t, image.Properties)
		}

		return true, nil
	})
}

func TestImagesListAllPages(t *testing.T) {
	client, err := clients.NewImageServiceV2Client()
	th.AssertNoErr(t, err)

	image, err := CreateEmptyImage(t, client)
	th.AssertNoErr(t, err)
	defer DeleteImage(t, client, image)

	listOpts := images.ListOpts{}

	allPages, err := images.List(client, listOpts).AllPages()
	th.AssertNoErr(t, err)

	allImages, err := images.ExtractImages(allPages)
	th.AssertNoErr(t, err)

	var found bool
	for _, i := range allImages {
		tools.PrintResource(t, i)
		tools.PrintResource(t, i.Properties)

		if i.Name == image.Name {
			found = true
		}
	}

	th.AssertEquals(t, found, true)
}

func TestImagesListByDate(t *testing.T) {
	client, err := clients.NewImageServiceV2Client()
	th.AssertNoErr(t, err)

	date := time.Date(2014, 1, 1, 1, 1, 1, 0, time.UTC)
	listOpts := images.ListOpts{
		Limit: 1,
		CreatedAtQuery: &images.ImageDateQuery{
			Date:   date,
			Filter: images.FilterGTE,
		},
	}

	allPages, err := images.List(client, listOpts).AllPages()
	th.AssertNoErr(t, err)

	allImages, err := images.ExtractImages(allPages)
	th.AssertNoErr(t, err)

	if len(allImages) == 0 {
		t.Fatalf("Query resulted in no results")
	}

	for _, image := range allImages {
		tools.PrintResource(t, image)
		tools.PrintResource(t, image.Properties)
	}

	date = time.Date(2049, 1, 1, 1, 1, 1, 0, time.UTC)
	listOpts = images.ListOpts{
		Limit: 1,
		CreatedAtQuery: &images.ImageDateQuery{
			Date:   date,
			Filter: images.FilterGTE,
		},
	}

	allPages, err = images.List(client, listOpts).AllPages()
	th.AssertNoErr(t, err)

	allImages, err = images.ExtractImages(allPages)
	th.AssertNoErr(t, err)

	if len(allImages) > 0 {
		t.Fatalf("Expected 0 images, got %d", len(allImages))
	}
}

func TestImagesFilter(t *testing.T) {
	client, err := clients.NewImageServiceV2Client()
	th.AssertNoErr(t, err)

	image, err := CreateEmptyImage(t, client)
	th.AssertNoErr(t, err)
	defer DeleteImage(t, client, image)

	listOpts := images.ListOpts{
		Tags:            []string{"foo", "bar"},
		ContainerFormat: "bare",
		DiskFormat:      "qcow2",
	}

	allPages, err := images.List(client, listOpts).AllPages()
	th.AssertNoErr(t, err)

	allImages, err := images.ExtractImages(allPages)
	th.AssertNoErr(t, err)

	if len(allImages) == 0 {
		t.Fatalf("Query resulted in no results")
	}
}
