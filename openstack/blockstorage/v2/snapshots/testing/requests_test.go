package testing

import (
	"testing"

	"time"

	"github.com/huaweicloud/golangsdk/openstack/blockstorage/v2/snapshots"
	th "github.com/huaweicloud/golangsdk/testhelper"
	"github.com/huaweicloud/golangsdk/testhelper/client"
)

func TestListAll(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	MockListResponse(t)

	allPages, err := snapshots.List(client.ServiceClient(), &snapshots.ListOpts{}).AllPages()
	th.AssertNoErr(t, err)
	actual, err := snapshots.ExtractSnapshots(allPages)
	th.AssertNoErr(t, err)

	expected := []snapshots.Snapshot{
		{
			ID:          "b836dc3d-4e10-4ea4-a34c-8f6b0460a583",
			Name:        "test01",
			Size:        1,
			Description: "",
			Status:      "available",
			VolumeID:    "ba5730ea-8621-4ae8-b702-ff0ffc12c209",
			Metadata:    map[string]string{},
			CreatedAt:   time.Date(2016, 2, 16, 16, 54, 14, 981520000, time.UTC),
		},
		{
			ID:          "83be494d-329e-4a78-8ac5-9af900f48b95",
			Name:        "test02",
			Size:        1,
			Description: "",
			Status:      "available",
			VolumeID:    "ba5730ea-8621-4ae8-b702-ff0ffc12c209",
			Metadata:    map[string]string{},
			CreatedAt:   time.Date(2016, 2, 16, 16, 54, 19, 475397000, time.UTC),
		},
	}

	th.CheckDeepEquals(t, expected, actual)

}

func TestDetail(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	MockDetailResponse(t)

	allPages, err := snapshots.Detail(client.ServiceClient(), &snapshots.ListOpts{}).AllPages()
	th.AssertNoErr(t, err)
	actual, err := snapshots.ExtractSnapshots(allPages)
	th.AssertNoErr(t, err)

	expected := []snapshots.Snapshot{
		{
			ID:          "6cd26877-3ca3-4f4e-ae2a-38cc3d6183fa",
			Name:        "name_xx2-snap",
			Size:        5,
			Description: "",
			Status:      "available",
			VolumeID:    "ae11e59c-bd56-434a-a00c-04757e1c066d",
			Metadata:    map[string]string{},
			ProjectID:   "d6c277ba8820452e83df36f33c9fa561",
			Progress:    "100%",
			CreatedAt:   time.Date(2013, 6, 19, 7, 15, 29, 0000000000, time.UTC),
		},
		{
			ID:          "b3253e26-5c37-48dd-8bf2-8795dd1e848f",
			Name:        "name_xx2-snap",
			Size:        5,
			Description: "",
			Status:      "available",
			VolumeID:    "ae11e59c-bd56-434a-a00c-04757e1c066d",
			Metadata:    map[string]string{},
			ProjectID:   "d6c277ba8820452e83df36f33c9fa561",
			Progress:    "100%",
			CreatedAt:   time.Date(2013, 6, 19, 9, 8, 8, 000000000, time.UTC),
		},
	}

	th.CheckDeepEquals(t, expected, actual)

}

func TestGet(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	MockGetResponse(t)

	v, err := snapshots.Get(client.ServiceClient(), "591ac654-26d8-41be-bb77-4f90699d2d41").Extract()
	th.AssertNoErr(t, err)

	th.AssertEquals(t, v.Name, "snap-001")
	th.AssertEquals(t, v.ID, "2bb856e1-b3d8-4432-a858-09e4ce939389")
}

func TestCreate(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	MockCreateResponse(t)

	options := &snapshots.CreateOpts{
		Name:        "snap-001",
		Description: "Daily backup",
		VolumeID:    "5aa119a8-d25b-45a7-8d1b-88e127885635",
		Force:       true,
		Metadata:    map[string]string{},
	}
	n, err := snapshots.Create(client.ServiceClient(), options).Extract()
	th.AssertNoErr(t, err)

	th.AssertEquals(t, n.Size, 1)
	th.AssertEquals(t, n.ID, "ffa9bc5e-1172-4021-acaf-cdcd78a9584d")
}

func TestDelete(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	MockDeleteResponse(t)

	res := snapshots.Delete(client.ServiceClient(), "d32019d3-bc6e-4319-9c1d-6722fc136a22")
	th.AssertNoErr(t, res.Err)
}

func TestUpdate(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	MockUpdateResponse(t)

	options := snapshots.UpdateOpts{Name: "name_xx3", Description: "hello"}
	v, err := snapshots.Update(client.ServiceClient(), "8dd7c486-8e9f-49fe-bceb-26aa7e312b66", options).Extract()
	th.AssertNoErr(t, err)
	th.CheckEquals(t, "snap-001", v.Name)
}

func TestRollback(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	MockRollbackResponse(t)

	options := snapshots.RollbackOpts{Name: "test-001", VolumeID: "5aa119a8-d25b-45a7-8d1b-88e127885635"}
	v, err := snapshots.Rollback(client.ServiceClient(), "8dd7c486-8e9f-49fe-bceb-26aa7e312b66", options).ExtractVolumeID()
	th.AssertNoErr(t, err)
	th.CheckEquals(t, "5aa119a8-d25b-45a7-8d1b-88e127885635", v)
}

func TestGetMetadata(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	MockGetMetadataResponse(t)

	v, err := snapshots.GetMetadata(client.ServiceClient(), "8dd7c486-8e9f-49fe-bceb-26aa7e312b66").ExtractMetadata()
	th.AssertNoErr(t, err)

	th.AssertEquals(t, v["key"], "val")
}

func TestCreateMetadata(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	MockCreateMetadataResponse(t)

	option := snapshots.MetadataOpts{
		Metadata: map[string]string{
			"key": "val",
		},
	}

	v, err := snapshots.CreateMetadata(client.ServiceClient(), "8dd7c486-8e9f-49fe-bceb-26aa7e312b66", &option).ExtractMetadata()
	th.AssertNoErr(t, err)

	th.AssertEquals(t, v["key"], "val")
}

func TestUpdateMetadata(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	MockUpdateMetadataResponse(t)

	option := snapshots.MetadataOpts{
		Metadata: map[string]string{
			"key": "val",
		},
	}

	v, err := snapshots.UpdateMetadata(client.ServiceClient(), "8dd7c486-8e9f-49fe-bceb-26aa7e312b66", &option).ExtractMetadata()
	th.AssertNoErr(t, err)

	th.AssertEquals(t, v["key"], "val")
}

func TestGetMetadataKey(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	MockGetMetadataKeyResponse(t)

	v, err := snapshots.GetMetadataKey(client.ServiceClient(), "8dd7c486-8e9f-49fe-bceb-26aa7e312b66", "key").ExtractMetadata()
	th.AssertNoErr(t, err)

	th.AssertEquals(t, v["key"], "val")
}

func TestUpdateMetadataKey(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	MockUpdateMetadataKeyResponse(t)

	option := snapshots.MetadataOpts{
		Metadata: map[string]string{
			"key": "val",
		},
	}

	v, err := snapshots.UpdateMetadataKey(client.ServiceClient(), "8dd7c486-8e9f-49fe-bceb-26aa7e312b66", "key", &option).ExtractMetadata()
	th.AssertNoErr(t, err)

	th.AssertEquals(t, v["key"], "val")
}

func TestDeleteMetadataKey(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	MockDeleteMetadataKeyResponse(t)

	res := snapshots.DeleteMetadataKey(client.ServiceClient(), "8dd7c486-8e9f-49fe-bceb-26aa7e312b66", "key")
	th.AssertNoErr(t, res.Err)
}
