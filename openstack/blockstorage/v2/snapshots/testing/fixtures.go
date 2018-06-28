package testing

import (
	"fmt"
	"net/http"
	"testing"

	th "github.com/huaweicloud/golangsdk/testhelper"
	fake "github.com/huaweicloud/golangsdk/testhelper/client"
)

func MockListResponse(t *testing.T) {
	th.Mux.HandleFunc("/snapshots", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Fprintf(w, `
{ 
    "snapshots": [ 
        { 
            "created_at": "2016-02-16T16:54:14.981520",  
            "description": null,  
            "id": "b836dc3d-4e10-4ea4-a34c-8f6b0460a583",  
            "metadata": { },  
            "name": "test01",  
            "size": 1,  
            "status": "available",  
            "volume_id": "ba5730ea-8621-4ae8-b702-ff0ffc12c209",  
            "updated_at": null 
        },  
        { 
            "created_at": "2016-02-16T16:54:19.475397",  
            "description": null,  
            "id": "83be494d-329e-4a78-8ac5-9af900f48b95",  
            "metadata": { },  
            "name": "test02",  
            "size": 1,  
            "status": "available",  
            "volume_id": "ba5730ea-8621-4ae8-b702-ff0ffc12c209",  
            "updated_at": null 
        }
    ],  
    "snapshots_links": null 
}
  `)
	})
}

func MockDetailResponse(t *testing.T) {
	th.Mux.HandleFunc("/snapshots/detail", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Fprintf(w, `
{ 
    "snapshots": [ 
        { 
            "status": "available",  
            "os-extended-snapshot-attributes:progress": "100%%",  
            "description": null,  
            "created_at": "2013-06-19T07:15:29.000000",  
            "metadata": { },  
            "volume_id": "ae11e59c-bd56-434a-a00c-04757e1c066d",  
            "os-extended-snapshot-attributes:project_id": "d6c277ba8820452e83df36f33c9fa561",  
            "size": 5,  
            "id": "6cd26877-3ca3-4f4e-ae2a-38cc3d6183fa",  
            "name": "name_xx2-snap",  
            "updated_at": null 
        },  
        { 
            "status": "available",  
            "os-extended-snapshot-attributes:progress": "100%%",  
            "description": null,  
            "created_at": "2013-06-19T09:08:08.000000",  
            "metadata": { },  
            "volume_id": "ae11e59c-bd56-434a-a00c-04757e1c066d",  
            "os-extended-snapshot-attributes:project_id": "d6c277ba8820452e83df36f33c9fa561",  
            "size": 5,  
            "id": "b3253e26-5c37-48dd-8bf2-8795dd1e848f",  
            "name": "name_xx2-snap",  
            "updated_at": null 
        } 
    ] 
}
  `)
	})
}

func MockGetResponse(t *testing.T) {
	th.Mux.HandleFunc("/snapshots/591ac654-26d8-41be-bb77-4f90699d2d41", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, `
{
    "snapshot": {
        "status": "available",
        "os-extended-snapshot-attributes:progress": "100%%",
        "description": "daily backup",
        "created_at": "2013-02-25T04:13:17.000000",
        "metadata": {},
        "volume_id": "5aa119a8-d25b-45a7-8d1b-88e127885635",
        "os-extended-snapshot-attributes:project_id": "0c2eba2c5af04d3f9e9d0d410b371fde",
        "size": 1,
        "id": "2bb856e1-b3d8-4432-a858-09e4ce939389",
        "name": "snap-001",
        "updated_at": null
    }
}
      `)
	})
}

func MockCreateResponse(t *testing.T) {
	th.Mux.HandleFunc("/snapshots", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		th.TestHeader(t, r, "Content-Type", "application/json")
		th.TestHeader(t, r, "Accept", "application/json")
		th.TestJSONRequest(t, r, `
{ 
    "snapshot": { 
        "name": "snap-001",  
        "description": "Daily backup",  
        "volume_id": "5aa119a8-d25b-45a7-8d1b-88e127885635",  
        "force": true,  
        "metadata": { } 
    } 
}

      `)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusAccepted)

		fmt.Fprintf(w, `
{ 
    "snapshot": { 
        "status": "creating",  
        "description": "Daily backup",  
        "created_at": "2013-02-25T03:56:53.081642",  
        "metadata": { },  
        "volume_id": "5aa119a8-d25b-45a7-8d1b-88e127885635",  
        "size": 1,  
        "id": "ffa9bc5e-1172-4021-acaf-cdcd78a9584d",  
        "name": "snap-001",  
        "updated_at": "2013-02-25T03:56:53.081642" 
    } 
}
    `)
	})
}

func MockDeleteResponse(t *testing.T) {
	th.Mux.HandleFunc("/snapshots/d32019d3-bc6e-4319-9c1d-6722fc136a22", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "DELETE")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		w.WriteHeader(http.StatusAccepted)
	})
}

func MockUpdateResponse(t *testing.T) {
	th.Mux.HandleFunc("/snapshots/8dd7c486-8e9f-49fe-bceb-26aa7e312b66", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "PUT")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		th.TestJSONRequest(t, r, `
{ 
    "snapshot": { 
        "name": "name_xx3",  
        "description": "hello" 
    } 
}
      `)
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, `
{ 
    "snapshot": { 
        "status": "creating",  
        "description": "Daily backup",  
        "created_at": "2013-02-25T03:56:53.081642",  
        "metadata": { },  
        "volume_id": "5aa119a8-d25b-45a7-8d1b-88e127885635",  
        "size": 1,  
        "id": "f9faf7df-fdc1-4093-9ef3-5cba06eef995",  
        "name": "snap-001",  
        "updated_at": "2013-02-25T03:56:53.081642" 
    } 
}
        `)
	})
}

func MockRollbackResponse(t *testing.T) {
	th.Mux.HandleFunc("/os-vendor-snapshots/8dd7c486-8e9f-49fe-bceb-26aa7e312b66/rollback", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		th.TestHeader(t, r, "Content-Type", "application/json")
		th.TestHeader(t, r, "Accept", "application/json")
		th.TestJSONRequest(t, r, `
{ 
    "rollback": { 
        "name": "test-001", 
        "volume_id": "5aa119a8-d25b-45a7-8d1b-88e127885635" 
    } 
}
      `)
		w.WriteHeader(http.StatusAccepted)
		fmt.Fprintf(w, `
{ 
    "rollback": { 
        "volume_id": "5aa119a8-d25b-45a7-8d1b-88e127885635" 
    } 
}
        `)
	})
}

func MockCreateMetadataResponse(t *testing.T) {
	th.Mux.HandleFunc("/snapshots/8dd7c486-8e9f-49fe-bceb-26aa7e312b66/metadata", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		th.TestHeader(t, r, "Content-Type", "application/json")
		th.TestHeader(t, r, "Accept", "application/json")
		th.TestJSONRequest(t, r, `
{ 
    "metadata": { 
        "key": "val"
    } 
}
      `)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusAccepted)

		fmt.Fprintf(w, `
{ 
    "metadata": { 
        "key": "val"
    } 
}
    `)
	})
}

func MockUpdateMetadataResponse(t *testing.T) {
	th.Mux.HandleFunc("/snapshots/8dd7c486-8e9f-49fe-bceb-26aa7e312b66/metadata", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "PUT")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		th.TestHeader(t, r, "Content-Type", "application/json")
		th.TestHeader(t, r, "Accept", "application/json")
		th.TestJSONRequest(t, r, `
{ 
    "metadata": { 
        "key": "val"
    } 
}

      `)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusAccepted)

		fmt.Fprintf(w, `
{ 
    "metadata": { 
        "key": "val"
    } 
}
    `)
	})
}

func MockGetMetadataResponse(t *testing.T) {
	th.Mux.HandleFunc("/snapshots/8dd7c486-8e9f-49fe-bceb-26aa7e312b66/metadata", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Fprintf(w, `
{ 
    "metadata": { 
       "key": "val"
    } 
}

    `)
	})
}

func MockGetMetadataKeyResponse(t *testing.T) {
	th.Mux.HandleFunc("/snapshots/8dd7c486-8e9f-49fe-bceb-26aa7e312b66/metadata/key", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Fprintf(w, `
{ 
    "metadata": { 
       "key": "val"
    } 
}

    `)
	})
}

func MockUpdateMetadataKeyResponse(t *testing.T) {
	th.Mux.HandleFunc("/snapshots/8dd7c486-8e9f-49fe-bceb-26aa7e312b66/metadata/key", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "PUT")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		th.TestHeader(t, r, "Content-Type", "application/json")
		th.TestHeader(t, r, "Accept", "application/json")
		th.TestJSONRequest(t, r, `
{ 
    "metadata": { 
        "key": "val"
    } 
}
      `)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusAccepted)

		fmt.Fprintf(w, `
{ 
    "metadata": { 
        "key": "val"
    } 
}
    `)
	})
}

func MockDeleteMetadataKeyResponse(t *testing.T) {
	th.Mux.HandleFunc("/snapshots/8dd7c486-8e9f-49fe-bceb-26aa7e312b66/metadata/key", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "DELETE")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		w.WriteHeader(http.StatusAccepted)
	})
}
