package testing

import (
	"testing"

	"github.com/huaweicloud/golangsdk/openstack/imageservice/v1/cloudimages"
	th "github.com/huaweicloud/golangsdk/testhelper"
	fakeclient "github.com/huaweicloud/golangsdk/testhelper/client"
)

func TestListImageTags(t *testing.T) {

	th.SetupHTTP()
	defer th.TeardownHTTP()

	HandleImageTagsListSuccessfully(t)

	r := cloudimages.ListImageTags(fakeclient.ServiceClient(), cloudimages.ListImageTagsOpts{})
	_, err := r.Extract()
	th.AssertNoErr(t, err)
}

func TestSetImageTag(t *testing.T) {

	th.SetupHTTP()
	defer th.TeardownHTTP()

	HandleImageTagSetSuccessfully(t)

	r := cloudimages.SetImageTag(fakeclient.ServiceClient(), cloudimages.SetImageTagOpts{ImageID: "imageID", Tag: "imageTag"})
	th.AssertNoErr(t, r.Err)
}

func TestImportImage(t *testing.T) {

	th.SetupHTTP()
	defer th.TeardownHTTP()

	HandleImageImportSuccessfully(t)

	r := cloudimages.ImportImage(fakeclient.ServiceClient(), "imageID", "imageURL")
	th.AssertNoErr(t, r.Err)
}

func TestExportImage(t *testing.T) {

	th.SetupHTTP()
	defer th.TeardownHTTP()

	HandleImageExportSuccessfully(t)

	r := cloudimages.ExportImage(fakeclient.ServiceClient(), "imageID", "bucketURL", "fileFormat")
	th.AssertNoErr(t, r.Err)
}

func TestCopyImage(t *testing.T) {

	th.SetupHTTP()
	defer th.TeardownHTTP()

	HandleImageCopySuccessfully(t)

	r := cloudimages.CopyImage(fakeclient.ServiceClient(), "imageID", cloudimages.CopyImageOpts{
		Name: "image_name",
	})
	_, err := r.Extract()
	th.AssertNoErr(t, err)
}

func TestAddImageMembers(t *testing.T) {

	th.SetupHTTP()
	defer th.TeardownHTTP()

	HandleImageMembersAddSuccessfully(t)

	r := cloudimages.AddImageMembers(fakeclient.ServiceClient(), cloudimages.AddImageMembersOpts{
		Images:   []string{"image1"},
		Projects: []string{"project1"},
	})
	_, err := r.Extract()
	th.AssertNoErr(t, err)
}

func TestUpdateImageMembers(t *testing.T) {

	th.SetupHTTP()
	defer th.TeardownHTTP()

	HandleImageMemberUpdateSuccessfully(t)

	r := cloudimages.UpdateImageMember(fakeclient.ServiceClient(), cloudimages.UpdateImageMemberOpts{
		Images:    []string{"image1"},
		ProjectID: "project1",
		Status:    "accepted",
	})
	_, err := r.Extract()
	th.AssertNoErr(t, err)
}

func TestDeleteImageMembers(t *testing.T) {

	th.SetupHTTP()
	defer th.TeardownHTTP()

	HandleImageMembersDeleteSuccessfully(t)

	r := cloudimages.DeleteImageMembers(fakeclient.ServiceClient(), cloudimages.DeleteImageMembersOpts{
		Images:   []string{"image1"},
		Projects: []string{"project1"},
	})
	_, err := r.Extract()
	th.AssertNoErr(t, err)
}
