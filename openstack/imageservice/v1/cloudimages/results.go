package cloudimages

import (
	"github.com/huaweicloud/golangsdk"
)

// ImageTagsResult represents the result of image tags query
type ImageTagsResult struct {
	golangsdk.Result
}

// ImageTags is the golang struct of ImageTagsResult
type ImageTags struct {
	Tags []string `json:"tags"`
}

// Extract extracts ImageTagsResult to struct
func (r ImageTagsResult) Extract() (*ImageTags, error) {
	var t *ImageTags
	err := r.ExtractInto(&t)
	return t, err
}

// SetImageTagResult is the result of image tag setting operation
type SetImageTagResult struct {
	golangsdk.Result
}

// ImportResult represents the result of the image importing
type ImportResult struct {
	golangsdk.Result
}

// ExportResult represents the result of the image exporting
type ExportResult struct {
	golangsdk.Result
}

// CopyResult represents the result of the image copying
type ImageCopyResult struct {
	golangsdk.Result
}

type asyncOperationJob struct {
	JobID string `json:"job_id"`
}

// ImageCopyTask is the image copy task description
type ImageCopyTask struct {
	asyncOperationJob
}

// Extract extracts ImageCopyResult to golangstruct
func (r ImageCopyResult) Extract() (*ImageCopyTask, error) {
	var t *ImageCopyTask
	err := r.ExtractInto(&t)
	return t, err
}

// AddImageMembersResult represents the result of image members adding operation
type AddImageMembersResult struct {
	golangsdk.Result
}

// AddImageMembersTask is the task description of image members adding operation
type AddImageMembersTask struct {
	asyncOperationJob
}

// Extract extracts AddImageMembersResult to golang struct
func (r AddImageMembersResult) Extract() (*AddImageMembersTask, error) {
	var t *AddImageMembersTask
	err := r.ExtractInto(&t)
	return t, err
}

// UpdateImageMemberResult represent the result of image members updating
// operation
type UpdateImageMemberResult struct {
	golangsdk.Result
}

// UpdateImageMemberTask is the task description of image member update
type UpdateImageMemberTask struct {
	asyncOperationJob
}

// Extract extracts UpdateImageMemberResult to golang struct
func (r UpdateImageMemberResult) Extract() (*UpdateImageMemberTask, error) {
	var t *UpdateImageMemberTask
	err := r.ExtractInto(&t)
	return t, err
}

// DeleteImageMembersResult represents the result of image members deleting
// operation
type DeleteImageMembersResult struct {
	golangsdk.Result
}

// DeleteImageMembersTask is the task description of image member deleting
type DeleteImageMembersTask struct {
	asyncOperationJob
}

// Extract extracts DeleteImageMembersResult to golang struct
func (r DeleteImageMembersResult) Extract() (*DeleteImageMembersTask, error) {
	var t *DeleteImageMembersTask
	err := r.ExtractInto(&t)
	return t, err
}

// CloudImagesQuota represents the quota of cloud images
type CloudImagesQuota struct {
	// Type is the resource type
	Type string `json:"type"`
	// Used is the amount of used resource
	Used int `json:"used"`
	// Quota is the amount of all resource
	Quota int `json:"quota"`
	// Min is the minimum of resource
	Min int `json:"min"`
	// Max is the maximum of resource
	Max int `json:"max"`
}

// CloudImagesQuotaResult represents the result of cloud image quota request
type CloudImagesQuotaResult struct {
	golangsdk.Result
}

// Extract extracts the result to CloudImagesQuota
func (r CloudImagesQuotaResult) Extract() (*CloudImagesQuota, error) {
	var s *CloudImagesQuota
	err := r.ExtractInto(&s)
	return s, err
}
