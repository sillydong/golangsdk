package testing

import (
	"testing"

	"github.com/huaweicloud/golangsdk"

	"time"

	"github.com/huaweicloud/golangsdk/openstack/blockstorage/v2/volumes"
	th "github.com/huaweicloud/golangsdk/testhelper"
	"github.com/huaweicloud/golangsdk/testhelper/client"
)

func TestListAll(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	MockListResponse(t)

	allPages, err := volumes.List(client.ServiceClient(), &volumes.ListOpts{}).AllPages()
	th.AssertNoErr(t, err)
	actual, err := volumes.ExtractVolumes(allPages)
	th.AssertNoErr(t, err)

	expected := []volumes.Volume{
		{
			ID:   "6b604cef-9bd8-4f5a-ae56-45839e6e1f0a",
			Name: "zjb_u25_test",
			Links: []golangsdk.Link{
				{
					Href: "https://volume.localdomain.com:8776/v2/dd14c6ac581f40059e27f5320b60bf2f/volumes/6b604cef-9bd8-4f5a-ae56-45839e6e1f0a",
					Rel:  "self",
				},
				{
					Href: "https://volume.localdomain.com:8776/dd14c6ac581f40059e27f5320b60bf2f/volumes/6b604cef-9bd8-4f5a-ae56-45839e6e1f0a",
					Rel:  "bookmark",
				},
			},
		},
		{
			ID:   "2bce4552-9a7d-48fa-8484-abbbf64b206e",
			Name: "zjb_u25_test",
			Links: []golangsdk.Link{
				{
					Href: "https://volume.localdomain.com:8776/v2/dd14c6ac581f40059e27f5320b60bf2f/volumes/2bce4552-9a7d-48fa-8484-abbbf64b206e",
					Rel:  "self",
				},
				{
					Href: "https://volume.localdomain.com:8776/dd14c6ac581f40059e27f5320b60bf2f/volumes/2bce4552-9a7d-48fa-8484-abbbf64b206e",
					Rel:  "bookmark",
				},
			},
		},
		{
			ID:   "3f1b98ec-a8b5-4e92-a727-88def62d5ad3",
			Name: "zjb_u25_test",
			Links: []golangsdk.Link{
				{
					Href: "https://volume.localdomain.com:8776/v2/dd14c6ac581f40059e27f5320b60bf2f/volumes/3f1b98ec-a8b5-4e92-a727-88def62d5ad3",
					Rel:  "self",
				},
				{
					Href: "https://volume.localdomain.com:8776/dd14c6ac581f40059e27f5320b60bf2f/volumes/3f1b98ec-a8b5-4e92-a727-88def62d5ad3",
					Rel:  "bookmark",
				},
			},
		},
	}

	th.CheckDeepEquals(t, expected, actual)

}

func TestDetail(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	MockDetailResponse(t)

	allPages, err := volumes.Detail(client.ServiceClient(), &volumes.ListOpts{}).AllPages()
	th.AssertNoErr(t, err)
	actual, err := volumes.ExtractVolumes(allPages)
	th.AssertNoErr(t, err)

	expected := []volumes.Volume{
		{
			ID:                 "b104b8db-170d-441b-897a-3c8ba9c5a214",
			Name:               "zjb_u25_test",
			Attachments:        []volumes.Attachment{},
			AvailabilityZone:   "xxx",
			Bootable:           "false",
			ConsistencyGroupID: "",
			CreatedAt:          time.Date(2016, 5, 25, 2, 42, 10, 856332000, time.UTC),
			UpdatedAt:          time.Date(2016, 5, 25, 2, 42, 22, 341984000, time.UTC),
			Description:        "",
			Encrypted:          false,
			Metadata: map[string]string{
				"__openstack_region_name": "pod01.xxx",
				"a":          "b",
				"quantityGB": "1",
				"volInfoUrl": "fusionstorage://172.30.64.10/0/FEFEEB07D3924CDEA93C612D4E16882D",
			},
			Multiattach:       false,
			ReplicationStatus: "disabled",
			Size:              1,
			SnapshotID:        "",
			SourceVolID:       "",
			Status:            "available",
			UserID:            "b0524e8342084ef5b74f158f78fc3049",
			VolumeType:        "SATA",
			Links: []golangsdk.Link{
				{
					Href: "https://volume.localdomain.com:8776/v2/dd14c6ac581f40059e27f5320b60bf2f/volumes/b104b8db-170d-441b-897a-3c8ba9c5a214",
					Rel:  "self",
				},
				{
					Href: "https://volume.localdomain.com:8776/dd14c6ac581f40059e27f5320b60bf2f/volumes/b104b8db-170d-441b-897a-3c8ba9c5a214",
					Rel:  "bookmark",
				},
			},
			Host:           "pod01.xxx#SATA",
			MigStat:        "",
			NameID:         "",
			TenentID:       "dd14c6ac581f40059e27f5320b60bf2f",
			DriverData:     "",
			ExtendedStatus: "",
		},
	}

	th.CheckDeepEquals(t, expected, actual)

}

func TestGet(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	MockGetResponse(t)

	v, err := volumes.Get(client.ServiceClient(), "591ac654-26d8-41be-bb77-4f90699d2d41").Extract()
	th.AssertNoErr(t, err)

	th.AssertEquals(t, v.Name, "restore_backup_0115efb3-678c-4a9e-bff6-d3cd278238b9")
	th.AssertEquals(t, v.ID, "591ac654-26d8-41be-bb77-4f90699d2d41")
}

func TestCreate(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	MockCreateResponse(t)

	options := &volumes.CreateOpts{
		Size:             40,
		Name:             "openapi_vol01",
		ImageID:          "027cf713-45a6-45f0-ac1b-0ccc57ac12e2",
		AvailabilityZone: "xxx",
		Description:      "create for api test",
		VolumeType:       "SATA",
		Metadata: map[string]string{
			"volume_owner": "openapi",
		},
		SchedulerHints: map[string]string{
			"dedicated_storage_id": "eddc1a3e-4145-45be-98d7-bf6f65af9767",
		},
		Shareable:   "false",
		Multiattach: false,
	}
	n, err := volumes.Create(client.ServiceClient(), options).Extract()
	th.AssertNoErr(t, err)

	th.AssertEquals(t, n.Size, 40)
	th.AssertEquals(t, n.ID, "8dd7c486-8e9f-49fe-bceb-26aa7e312b66")
}

func TestDelete(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	MockDeleteResponse(t)

	res := volumes.Delete(client.ServiceClient(), "d32019d3-bc6e-4319-9c1d-6722fc136a22")
	th.AssertNoErr(t, res.Err)
}

func TestUpdate(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	MockUpdateResponse(t)

	options := volumes.UpdateOpts{Name: "openapi_vol01", Description: "create for api test"}
	v, err := volumes.Update(client.ServiceClient(), "8dd7c486-8e9f-49fe-bceb-26aa7e312b66", options).Extract()
	th.AssertNoErr(t, err)
	th.CheckEquals(t, "openapi_vol01", v.Name)
}

func TestGetMetadata(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	MockGetMetadataResponse(t)

	v, err := volumes.GetMetadata(client.ServiceClient(), "8dd7c486-8e9f-49fe-bceb-26aa7e312b66").ExtractMetadata()
	th.AssertNoErr(t, err)

	th.AssertEquals(t, v["key"], "val")
}

func TestCreateMetadata(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	MockCreateMetadataResponse(t)

	option := volumes.MetadataOpts{
		Metadata: map[string]string{
			"key": "val",
		},
	}

	v, err := volumes.CreateMetadata(client.ServiceClient(), "8dd7c486-8e9f-49fe-bceb-26aa7e312b66", &option).ExtractMetadata()
	th.AssertNoErr(t, err)

	th.AssertEquals(t, v["key"], "val")
}

func TestUpdateMetadata(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	MockUpdateMetadataResponse(t)

	option := volumes.MetadataOpts{
		Metadata: map[string]string{
			"key": "val",
		},
	}

	v, err := volumes.UpdateMetadata(client.ServiceClient(), "8dd7c486-8e9f-49fe-bceb-26aa7e312b66", &option).ExtractMetadata()
	th.AssertNoErr(t, err)

	th.AssertEquals(t, v["key"], "val")
}

func TestGetMetadataKey(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	MockGetMetadataKeyResponse(t)

	v, err := volumes.GetMetadataKey(client.ServiceClient(), "8dd7c486-8e9f-49fe-bceb-26aa7e312b66", "key").ExtractMetadata()
	th.AssertNoErr(t, err)

	th.AssertEquals(t, v["key"], "val")
}

func TestUpdateMetadataKey(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	MockUpdateMetadataKeyResponse(t)

	option := volumes.MetadataOpts{
		Metadata: map[string]string{
			"key": "val",
		},
	}

	v, err := volumes.UpdateMetadataKey(client.ServiceClient(), "8dd7c486-8e9f-49fe-bceb-26aa7e312b66", "key", &option).ExtractMetadata()
	th.AssertNoErr(t, err)

	th.AssertEquals(t, v["key"], "val")
}

func TestDeleteMetadataKey(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	MockDeleteMetadataKeyResponse(t)

	res := volumes.DeleteMetadataKey(client.ServiceClient(), "8dd7c486-8e9f-49fe-bceb-26aa7e312b66", "key")
	th.AssertNoErr(t, res.Err)
}

func TestExtendSize(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	MockExtendSizeResponse(t)

	options := &volumes.ExtendSizeOpts{
		NewSize: 3,
	}

	err := volumes.ExtendSize(client.ServiceClient(), "8dd7c486-8e9f-49fe-bceb-26aa7e312b66", options).ExtractErr()
	th.AssertNoErr(t, err)
}

func TestSetBootable(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	MockSetBootableResponse(t)

	options := &volumes.SetBootableOpts{
		Bootable: true,
	}

	err := volumes.SetBootable(client.ServiceClient(), "8dd7c486-8e9f-49fe-bceb-26aa7e312b66", options).ExtractErr()
	th.AssertNoErr(t, err)
}

func TestSetReadOnly(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	MockSetReadOnlyResponse(t)

	options := &volumes.SetReadOnlyOpts{
		ReadOnly: true,
	}

	err := volumes.SetReadOnly(client.ServiceClient(), "8dd7c486-8e9f-49fe-bceb-26aa7e312b66", options).ExtractErr()
	th.AssertNoErr(t, err)
}
