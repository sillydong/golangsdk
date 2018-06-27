package testing

import (
	"fmt"
	"net/http"
	"testing"

	th "github.com/huaweicloud/golangsdk/testhelper"
	fake "github.com/huaweicloud/golangsdk/testhelper/client"
)

func MockListResponse(t *testing.T) {
	th.Mux.HandleFunc("/volumes", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Fprintf(w, `
{
    "volumes": [
        {
            "id": "6b604cef-9bd8-4f5a-ae56-45839e6e1f0a",
            "links": [
                {
                    "href": "https://volume.localdomain.com:8776/v2/dd14c6ac581f40059e27f5320b60bf2f/volumes/6b604cef-9bd8-4f5a-ae56-45839e6e1f0a",
                    "rel": "self"
                },
                {
                    "href": "https://volume.localdomain.com:8776/dd14c6ac581f40059e27f5320b60bf2f/volumes/6b604cef-9bd8-4f5a-ae56-45839e6e1f0a",
                    "rel": "bookmark"
                }
            ],
            "name": "zjb_u25_test"
        },
        {
            "id": "2bce4552-9a7d-48fa-8484-abbbf64b206e",
            "links": [
                {
                    "href": "https://volume.localdomain.com:8776/v2/dd14c6ac581f40059e27f5320b60bf2f/volumes/2bce4552-9a7d-48fa-8484-abbbf64b206e",
                    "rel": "self"
                },
                {
                    "href": "https://volume.localdomain.com:8776/dd14c6ac581f40059e27f5320b60bf2f/volumes/2bce4552-9a7d-48fa-8484-abbbf64b206e",
                    "rel": "bookmark"
                }
            ],
            "name": "zjb_u25_test"
        },
        {
            "id": "3f1b98ec-a8b5-4e92-a727-88def62d5ad3",
            "links": [
                {
                    "href": "https://volume.localdomain.com:8776/v2/dd14c6ac581f40059e27f5320b60bf2f/volumes/3f1b98ec-a8b5-4e92-a727-88def62d5ad3",
                    "rel": "self"
                },
                {
                    "href": "https://volume.localdomain.com:8776/dd14c6ac581f40059e27f5320b60bf2f/volumes/3f1b98ec-a8b5-4e92-a727-88def62d5ad3",
                    "rel": "bookmark"
                }
            ],
            "name": "zjb_u25_test"
        }
    ],
    "volumes_links": [
        {
            "href": "https://volume.localdomain.com:8776/v2/dd14c6ac581f40059e27f5320b60bf2f/volumes?limit=3&marker=3f1b98ec-a8b5-4e92-a727-88def62d5ad3",
            "rel": "next"
        }
    ]
}
  `)
	})
}

func MockDetailResponse(t *testing.T) {
	th.Mux.HandleFunc("/volumes/detail", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Fprintf(w, `
{
    "volumes": [
        {
            "attachments": [],
            "availability_zone": "xxx",
            "bootable": "false",
            "consistencygroup_id": null,
            "created_at": "2016-05-25T02:42:10.856332",
            "description": null,
            "encrypted": false,
            "id": "b104b8db-170d-441b-897a-3c8ba9c5a214",
            "links": [
                {
                    "href": "https://volume.localdomain.com:8776/v2/dd14c6ac581f40059e27f5320b60bf2f/volumes/b104b8db-170d-441b-897a-3c8ba9c5a214",
                    "rel": "self"
                },
                {
                    "href": "https://volume.localdomain.com:8776/dd14c6ac581f40059e27f5320b60bf2f/volumes/b104b8db-170d-441b-897a-3c8ba9c5a214",
                    "rel": "bookmark"
                }
            ],
            "metadata": {
                "__openstack_region_name": "pod01.xxx",
                "a": "b",
                "quantityGB": "1",
                "volInfoUrl": "fusionstorage://172.30.64.10/0/FEFEEB07D3924CDEA93C612D4E16882D"
            },
            "name": "zjb_u25_test",
            "os-vol-host-attr:host": "pod01.xxx#SATA",
            "volume_image_metadata": {},
            "os-vol-mig-status-attr:migstat": null,
            "os-vol-mig-status-attr:name_id": null,
            "os-vol-tenant-attr:tenant_id": "dd14c6ac581f40059e27f5320b60bf2f",
            "os-volume-replication:driver_data": null,
            "os-volume-replication:extended_status": null,
            "replication_status": "disabled",
            "shareable": false,
            "multiattach": false,
            "size": 1,
            "snapshot_id": null,
            "source_volid": null,
            "status": "available",
            "updated_at": "2016-05-25T02:42:22.341984",
            "user_id": "b0524e8342084ef5b74f158f78fc3049",
            "volume_type": "SATA"
        }
    ],
    "volumes_links": [
        {
            "href": "https://volume.localdomain.com:8776/v2/dd14c6ac581f40059e27f5320b60bf2f/volumes/detail?limit=1&marker=b104b8db-170d-441b-897a-3c8ba9c5a214",
            "rel": "next"
        }
    ]
}
  `)
	})
}

func MockGetResponse(t *testing.T) {
	th.Mux.HandleFunc("/volumes/591ac654-26d8-41be-bb77-4f90699d2d41", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, `
{
    "volume": {
        "attachments": [],
        "links": [
            {
                "href": "https://volume.az0.dc1.domainname.com/v2/40acc331ac784f34842ba4f08ff2be48/volumes/591ac654-26d8-41be-bb77-4f90699d2d41",
                "rel": "self"
            },
            {
                "href": "https://volume.az0.dc1.domainname.com/40acc331ac784f34842ba4f08ff2be48/volumes/591ac654-26d8-41be-bb77-4f90699d2d41",
                "rel": "bookmark"
            }
        ],
        "availability_zone": "az1.dc1",
        "os-vol-host-attr:host": "az1.dc1#SSD",
        "encrypted": false,
        "multiattach": true,
        "updated_at": "2016-02-03T02:19:29.895237",
        "os-volume-replication:extended_status": null,
        "replication_status": "disabled",
        "snapshot_id": null,
        "id": "591ac654-26d8-41be-bb77-4f90699d2d41",
        "size": 40,
        "user_id": "fd03ee73295e45478d88e15263d2ee4e",
        "os-vol-tenant-attr:tenant_id": "40acc331ac784f34842ba4f08ff2be48",
        "volume_image_metadata": null,
        "os-vol-mig-status-attr:migstat": null,
        "metadata": {
            "__openstack_region_name": "az1.dc1",
            "quantityGB": "40"
        },
        "status": "error_restoring",
        "description": "auto-created_from_restore_from_backup",
        "os-volume-replication:driver_data": null,
        "source_volid": null,
        "consistencygroup_id": null,
        "os-vol-mig-status-attr:name_id": null,
        "name": "restore_backup_0115efb3-678c-4a9e-bff6-d3cd278238b9",
        "bootable": "false",
        "created_at": "2016-02-03T02:19:11.723797",
        "volume_type": null,
        "shareable": false
    }
}
      `)
	})
}

func MockCreateResponse(t *testing.T) {
	th.Mux.HandleFunc("/volumes", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		th.TestHeader(t, r, "Content-Type", "application/json")
		th.TestHeader(t, r, "Accept", "application/json")
		th.TestJSONRequest(t, r, `
{
    "volume": {
        "name": "openapi_vol01",
        "imageRef": "027cf713-45a6-45f0-ac1b-0ccc57ac12e2",
        "availability_zone": "xxx",
        "description": "create for api test",
        "volume_type": "SATA",
        "metadata": {
            "volume_owner": "openapi"
        },
        "OS-SCH-HNT:scheduler_hints": {
            "dedicated_storage_id": "eddc1a3e-4145-45be-98d7-bf6f65af9767"
        },
        "shareable": "false",
        "multiattach": false,
        "size": 40
    }
}
      `)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusAccepted)

		fmt.Fprintf(w, `
{
    "volume": {
        "attachments": [],
        "availability_zone": "xxx",
        "bootable": "false",
        "consistencygroup_id": null,
        "created_at": "2016-05-25T02:38:40.392463",
        "description": "create for api test",
        "encrypted": false,
        "id": "8dd7c486-8e9f-49fe-bceb-26aa7e312b66",
        "links": [
            {
                "href": "https://volume.localdomain.com:8776/v2/5dd0b0056f3d47b6ab4121667d35621a/volumes/8dd7c486-8e9f-49fe-bceb-26aa7e312b66",
                "rel": "self"
            },
            {
                "href": "https://volume.localdomain.com:8776/5dd0b0056f3d47b6ab4121667d35621a/volumes/8dd7c486-8e9f-49fe-bceb-26aa7e312b66",
                "rel": "bookmark"
            }
        ],
        "metadata": {
            "volume_owner": "openapi"
        },
        "name": "openapi_vol01",
        "replication_status": "disabled",
        "shareable": false,
        "multiattach": false,
        "size": 40,
        "snapshot_id": null,
        "source_volid": null,
        "status": "creating",
        "updated_at": null,
        "user_id": "39f6696ae23740708d0f358a253c2637",
        "volume_type": "SATA"
    }
}
    `)
	})
}

func MockDeleteResponse(t *testing.T) {
	th.Mux.HandleFunc("/volumes/d32019d3-bc6e-4319-9c1d-6722fc136a22", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "DELETE")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		w.WriteHeader(http.StatusAccepted)
	})
}

func MockUpdateResponse(t *testing.T) {
	th.Mux.HandleFunc("/volumes/8dd7c486-8e9f-49fe-bceb-26aa7e312b66", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "PUT")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		th.TestJSONRequest(t, r, `
{ 
    "volume": { 
        "name": "openapi_vol01",  
        "description": "create for api test" 
    } 
}
      `)
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, `
{
    "volume": {
        "attachments": [],
        "availability_zone": "xxx",
        "bootable": "false",
        "consistencygroup_id": null,
        "created_at": "2016-05-25T02:38:40.392463",
        "description": "create for api test",
        "encrypted": false,
        "id": "8dd7c486-8e9f-49fe-bceb-26aa7e312b66",
        "links": [
            {
                "href": "https://volume.localdomain.com:8776/v2/5dd0b0056f3d47b6ab4121667d35621a/volumes/8dd7c486-8e9f-49fe-bceb-26aa7e312b66",
                "rel": "self"
            },
            {
                "href": "https://volume.localdomain.com:8776/5dd0b0056f3d47b6ab4121667d35621a/volumes/8dd7c486-8e9f-49fe-bceb-26aa7e312b66",
                "rel": "bookmark"
            }
        ],
        "metadata": {
            "volume_owner": "openapi"
        },
        "name": "openapi_vol01",
        "replication_status": "disabled",
        "shareable": false,
        "multiattach": false,
        "size": 40,
        "snapshot_id": null,
        "source_volid": null,
        "status": "creating",
        "updated_at": null,
        "user_id": "39f6696ae23740708d0f358a253c2637",
        "volume_type": "SATA"
    }
}
        `)
	})
}

func MockCreateMetadataResponse(t *testing.T) {
	th.Mux.HandleFunc("/volumes/8dd7c486-8e9f-49fe-bceb-26aa7e312b66/metadata", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		th.TestHeader(t, r, "Content-Type", "application/json")
		th.TestHeader(t, r, "Accept", "application/json")
		th.TestJSONRequest(t, r, `
{
    "meta": {
        "key": "val"
    }
}
      `)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusAccepted)

		fmt.Fprintf(w, `
{
    "meta": {
        "key": "val"
    }
}
    `)
	})
}

func MockUpdateMetadataResponse(t *testing.T) {
	th.Mux.HandleFunc("/volumes/8dd7c486-8e9f-49fe-bceb-26aa7e312b66/metadata", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "PUT")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		th.TestHeader(t, r, "Content-Type", "application/json")
		th.TestHeader(t, r, "Accept", "application/json")
		th.TestJSONRequest(t, r, `
{
    "meta": {
        "key": "val"
    }
}
      `)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusAccepted)

		fmt.Fprintf(w, `
{
    "meta": {
        "key": "val"
    }
}
    `)
	})
}

func MockGetMetadataResponse(t *testing.T) {
	th.Mux.HandleFunc("/volumes/8dd7c486-8e9f-49fe-bceb-26aa7e312b66/metadata", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Fprintf(w, `
{
    "meta": {
        "key": "val"
    }
}
    `)
	})
}

func MockGetMetadataKeyResponse(t *testing.T) {
	th.Mux.HandleFunc("/volumes/8dd7c486-8e9f-49fe-bceb-26aa7e312b66/metadata/key", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Fprintf(w, `
{
    "meta": {
        "key": "val"
    }
}
    `)
	})
}

func MockUpdateMetadataKeyResponse(t *testing.T) {
	th.Mux.HandleFunc("/volumes/8dd7c486-8e9f-49fe-bceb-26aa7e312b66/metadata/key", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "PUT")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		th.TestHeader(t, r, "Content-Type", "application/json")
		th.TestHeader(t, r, "Accept", "application/json")
		th.TestJSONRequest(t, r, `
{
    "meta": {
        "key": "val"
    }
}
      `)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusAccepted)

		fmt.Fprintf(w, `
{
    "meta": {
        "key": "val"
    }
}
    `)
	})
}

func MockDeleteMetadataKeyResponse(t *testing.T) {
	th.Mux.HandleFunc("/volumes/8dd7c486-8e9f-49fe-bceb-26aa7e312b66/metadata/key", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "DELETE")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		w.WriteHeader(http.StatusAccepted)
	})
}

func MockExtendSizeResponse(t *testing.T) {
	th.Mux.HandleFunc("/volumes/8dd7c486-8e9f-49fe-bceb-26aa7e312b66/action",
		func(w http.ResponseWriter, r *http.Request) {
			th.TestMethod(t, r, "POST")
			th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
			th.TestHeader(t, r, "Content-Type", "application/json")
			th.TestHeader(t, r, "Accept", "application/json")
			th.TestJSONRequest(t, r, `
{
    "os-extend":
    {
        "new_size": 3
    }
}
          `)

			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusAccepted)

			fmt.Fprintf(w, `{}`)
		})
}

func MockSetBootableResponse(t *testing.T) {
	th.Mux.HandleFunc("/volumes/8dd7c486-8e9f-49fe-bceb-26aa7e312b66/action",
		func(w http.ResponseWriter, r *http.Request) {
			th.TestMethod(t, r, "POST")
			th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
			th.TestHeader(t, r, "Content-Type", "application/json")
			th.TestHeader(t, r, "Accept", "application/json")
			th.TestJSONRequest(t, r, `
{
    "os-set_bootable":
    {
        "bootable": true
    }
}
          `)

			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusAccepted)

			fmt.Fprintf(w, `{}`)
		})
}

func MockSetReadOnlyResponse(t *testing.T) {
	th.Mux.HandleFunc("/volumes/8dd7c486-8e9f-49fe-bceb-26aa7e312b66/action",
		func(w http.ResponseWriter, r *http.Request) {
			th.TestMethod(t, r, "POST")
			th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
			th.TestHeader(t, r, "Content-Type", "application/json")
			th.TestHeader(t, r, "Accept", "application/json")
			th.TestJSONRequest(t, r, `
{
    "os-update_readonly_flag":
    {
        "readonly": true
    }
}
          `)

			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusAccepted)

			fmt.Fprintf(w, `{}`)
		})
}
