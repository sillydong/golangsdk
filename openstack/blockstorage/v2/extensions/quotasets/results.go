package quotasets

import (
	"github.com/huaweicloud/golangsdk"
)

// QuotaSet is a set of operational limits that allow for control of block
// storage usage.

// QuotaUsageSet represents details of both operational limits of block
// storage resources and the current usage of those resources.
type QuotaUsageSet struct {
	// ID is the project ID associated with this QuotaUsageSet.
	ID string `json:"id"`

	// Volumes is the volume usage information for this project, including
	// in_use, limit, reserved and allocated attributes. Note: allocated
	// attribute is available only when nested quota is enabled.
	Volumes QuotaUsage `json:"volumes"`

	// Snapshots is the snapshot usage information for this project, including
	// in_use, limit, reserved and allocated attributes. Note: allocated
	// attribute is available only when nested quota is enabled.
	Snapshots QuotaUsage `json:"snapshots"`

	// Gigabytes is the size (GB) usage information of volumes and snapshots
	// for this project, including in_use, limit, reserved and allocated
	// attributes. Note: allocated attribute is available only when nested
	// quota is enabled.
	Gigabytes QuotaUsage `json:"gigabytes"`

	VolumesSAS   QuotaUsage `json:"volumes_SAS"`
	SnapshotsSAS QuotaUsage `json:"snapshots_SAS"`
	GigabytesSAS QuotaUsage `json:"gigabytes_SAS"`

	VolumesSATA   QuotaUsage `json:"volumes_SATA"`
	SnapshotsSATA QuotaUsage `json:"snapshots_SATA"`
	GigabytesSATA QuotaUsage `json:"gigabytes_SATA"`

	VolumesSSD   QuotaUsage `json:"volumes_SSD"`
	SnapshotsSSD QuotaUsage `json:"snapshots_SSD"`
	GigabytesSSD QuotaUsage `json:"gigabytes_SSD"`

	// Backups is the backup usage information for this project, including
	// in_use, limit, reserved and allocated attributes. Note: allocated
	// attribute is available only when nested quota is enabled.
	Backups QuotaUsage `json:"backups"`

	// BackupGigabytes is the size (GB) usage information of backup for this
	// project, including in_use, limit, reserved and allocated attributes.
	// Note: allocated attribute is available only when nested quota is
	// enabled.
	BackupGigabytes QuotaUsage `json:"backup_gigabytes"`
}

// QuotaUsage is a set of details about a single operational limit that allows
// for control of block storage usage.
type QuotaUsage struct {
	// InUse is the current number of provisioned resources of the given type.
	InUse int `json:"in_use"`

	// Allocated is the current number of resources of a given type allocated
	// for use.  It is only available when nested quota is enabled.
	Allocated int `json:"allocated"`

	// Reserved is a transitional state when a claim against quota has been made
	// but the resource is not yet fully online.
	Reserved int `json:"reserved"`

	// Limit is the maximum number of a given resource that can be
	// allocated/provisioned.  This is what "quota" usually refers to.
	Limit int `json:"limit"`
}

type quotaUsageResult struct {
	golangsdk.Result
}

// GetUsageResult is the response from a Get operation. Call its Extract
// method to interpret it as a QuotaSet.
type GetUsageResult struct {
	quotaUsageResult
}

// Extract is a method that attempts to interpret any QuotaUsageSet resource
// response as a set of QuotaUsageSet structs.
func (r quotaUsageResult) Extract() (QuotaUsageSet, error) {
	var s struct {
		QuotaUsageSet QuotaUsageSet `json:"quota_set"`
	}
	err := r.ExtractInto(&s)
	return s.QuotaUsageSet, err
}
