package testing

import (
	"fmt"
	"net/http"
	"testing"

	th "github.com/huaweicloud/golangsdk/testhelper"
)

// HandleImageTagsListSuccessfully test setup
func HandleImageTagsListSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/cloudimages/tags", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		w.Header().Add("Content-Type", "application/json")

		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, `{
			"tags": [
				"tag1",
				"tag2"
			]
		}`)
	})
}

// HandleImageTagSetSuccessfully test setup
func HandleImageTagSetSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/cloudimages/tags", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "PUT")

		w.WriteHeader(http.StatusNoContent)
	})
}

// HandleImageImportSuccessfully test setup
func HandleImageImportSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/cloudimages/imageID/upload", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "PUT")
		th.TestJSONRequest(t, r, `{"image_url":"imageURL"}`)

		w.WriteHeader(http.StatusOK)
	})
}

// HandleImageExportSuccessfully test setup
func HandleImageExportSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/cloudimages/imageID/file", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestJSONRequest(t, r, `{"bucket_url":"bucketURL", "file_format":"fileFormat"}`)

		w.WriteHeader(http.StatusOK)
	})
}

func HandleImageCopySuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/cloudimages/imageID/copy", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestJSONRequest(t, r, `{"name": "image_name"}`)

		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, `{
			"job_id": "123321"
		}`)
	})
}

func HandleImageMembersAddSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/cloudimages/members", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestJSONRequest(t, r, `{"images":["image1"], "projects":["project1"]}`)

		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, `{
			"job_id": "123321"
		}`)
	})
}

func HandleImageMemberUpdateSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/cloudimages/members", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "PUT")
		th.TestJSONRequest(t, r, `{"images":["image1"], "project_id":"project1", "status": "accepted"}`)

		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, `{
			"job_id": "123321"
		}`)
	})
}

func HandleImageMembersDeleteSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/cloudimages/members", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "DELETE")
		th.TestJSONRequest(t, r, `{"images":["image1"], "projects":["project1"]}`)

		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, `{
			"job_id": "123321"
		}`)
	})
}
