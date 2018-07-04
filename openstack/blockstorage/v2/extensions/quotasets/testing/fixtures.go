package testing

import (
	"fmt"
	"net/http"
	"testing"

	th "github.com/huaweicloud/golangsdk/testhelper"
	fake "github.com/huaweicloud/golangsdk/testhelper/client"
)

func MockGetUsageResponse(t *testing.T) {
	th.Mux.HandleFunc("/os-quota-sets/"+FirstTenantID, func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Fprintf(w, `
{ 
    "quota_set": { 
        "gigabytes_SAS": { 
            "reserved": 0,  
            "limit": -1,  
            "in_use": 21 
        },  
        "volumes_SATA": { 
            "reserved": 0,  
            "limit": -1,  
            "in_use": 8 
        },  
        "gigabytes": { 
            "reserved": 0,  
            "limit": 42790,  
            "in_use": 2792 
        },  
        "backup_gigabytes": { 
            "reserved": 0,  
            "limit": 5120,  
            "in_use": 51 
        },  
        "snapshots_SAS": { 
            "reserved": 0,  
            "limit": -1,  
            "in_use": 0 
        },  
        "volumes_SSD": { 
            "reserved": 0,  
            "limit": -1,  
            "in_use": 28 
        },  
        "snapshots": { 
            "reserved": 0,  
            "limit": 10,  
            "in_use": 6 
        },  
        "id": "cd631140887d4b6e9c786b67a6dd4c02",  
        "volumes_SAS": { 
            "reserved": 0,  
            "limit": -1,  
            "in_use": 2 
        },  
        "snapshots_SSD": { 
            "reserved": 0,  
            "limit": -1,  
            "in_use": 0 
        },  
        "volumes": { 
            "reserved": 0,  
            "limit": -1,  
            "in_use": 108 
        },  
        "gigabytes_SATA": { 
            "reserved": 0,  
            "limit": -1,  
            "in_use": 168 
        },  
        "backups": { 
            "reserved": 0,  
            "limit": 100,  
            "in_use": 10 
        },  
        "gigabytes_SSD": { 
            "reserved": 0,  
            "limit": -1,  
            "in_use": 1085 
        },  
        "snapshots_SATA": { 
            "reserved": 0,  
            "limit": -1,  
            "in_use": 0 
        } 
    } 
}

  `)
	})
}

const FirstTenantID = "555544443333222211110000ffffeeee"
