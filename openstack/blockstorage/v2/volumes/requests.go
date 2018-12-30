package volumes

import (
	"github.com/huaweicloud/golangsdk"
	"github.com/huaweicloud/golangsdk/pagination"
)

// CreateOptsBuilder allows extensions to add additional parameters to the
// Create request.
type CreateOptsBuilder interface {
	ToVolumeCreateMap() (map[string]interface{}, error)
}

// CreateOpts contains options for creating a Volume. This object is passed to
// the volumes.Create function. For more information about these parameters,
// see the Volume object.
type CreateOpts struct {
	// The size of the volume, in GB
	Size int `json:"size" required:"true"`
	// The availability zone
	AvailabilityZone string `json:"availability_zone" required:"true"`
	// ConsistencyGroupID is the ID of a consistency group
	ConsistencyGroupID string `json:"consistencygroup_id,omitempty"`
	// The volume description
	Description string `json:"description,omitempty"`
	// One or more metadata key and value pairs to associate with the volume
	Metadata map[string]string `json:"metadata,omitempty"`
	// The volume name
	Name string `json:"name,omitempty"`
	// the ID of the existing volume snapshot
	SnapshotID string `json:"snapshot_id,omitempty"`
	// SourceReplica is a UUID of an existing volume to replicate with
	SourceReplica string `json:"source_replica,omitempty"`
	// the ID of the existing volume
	SourceVolID string `json:"source_volid,omitempty"`
	// The ID of the image from which you want to create the volume.
	// Required to create a bootable volume.
	ImageID string `json:"imageRef,omitempty"`
	// The associated volume type
	VolumeType string `json:"volume_type,omitempty"`
	// Flag defines if the volume is shareable, deprecated
	Shareable string `json:"shareable,omitempty"`
	// Flag defines if the volume is shareable
	Multiattach bool `json:"multiattach"`
	// Schedule parameters, currently supports dedicated_storage_id
	SchedulerHints map[string]string `json:"OS-SCH-HNT:scheduler_hints,omitempty"`
}

// ToVolumeCreateMap assembles a request body based on the contents of a
// CreateOpts.
func (opts CreateOpts) ToVolumeCreateMap() (map[string]interface{}, error) {
	return golangsdk.BuildRequestBody(opts, "volume")
}

// Create will create a new Volume based on the values in CreateOpts. To extract
// the Volume object from the response, call the Extract method on the
// CreateResult.
func Create(client *golangsdk.ServiceClient, opts CreateOptsBuilder) (r CreateResult) {
	b, err := opts.ToVolumeCreateMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = client.Post(createURL(client), b, &r.Body, &golangsdk.RequestOpts{
		OkCodes: []int{202},
	})
	return
}

// Delete will delete the existing Volume with the provided ID.
func Delete(client *golangsdk.ServiceClient, id string) (r DeleteResult) {
	_, r.Err = client.Delete(deleteURL(client, id), nil)
	return
}

// Get retrieves the Volume with the provided ID. To extract the Volume object
// from the response, call the Extract method on the GetResult.
func Get(client *golangsdk.ServiceClient, id string) (r GetResult) {
	_, r.Err = client.Get(getURL(client, id), &r.Body, nil)
	return
}

// ListOptsBuilder allows extensions to add additional parameters to the List
// request.
type ListOptsBuilder interface {
	ToVolumeListQuery() (string, error)
}

// ListOpts holds options for listing Volumes. It is passed to the volumes.List
// function.
type ListOpts struct {
	// Metadata will filter results based on specified metadata.
	Metadata string `q:"metadata"`

	// Name will filter by the specified volume name.
	Name string `q:"name"`

	// Status will filter by the specified status.
	Status string `q:"status"`

	// TenantID will filter by a specific tenant/project ID.
	TenantID string `q:"project_id"`

	// Limit of pagination
	Limit int `q:"limit"`

	// Sort result by key
	SortKey string `q:"sort_key"`

	// Sort direction, desc or asc
	SortDir string `q:"sort_dir"`

	// Offset of pagination
	Offset int `q:"offset"`

	// Avalibility zone information
	AvailabilityZone string `q:"availablity_zone"`

	// Query update time of could volume
	ChangeSince string `q:"change_since"`

	// Last record id of last page
	Marker string `q:"marker"`
}

// ToVolumeListQuery formats a ListOpts into a query string.
func (opts ListOpts) ToVolumeListQuery() (string, error) {
	q, err := golangsdk.BuildQueryString(opts)
	return q.String(), err
}

// List returns Volumes optionally limited by the conditions provided in ListOpts.
func List(client *golangsdk.ServiceClient, opts ListOptsBuilder) pagination.Pager {
	url := listURL(client)
	if opts != nil {
		query, err := opts.ToVolumeListQuery()
		if err != nil {
			return pagination.Pager{Err: err}
		}
		url += query
	}

	return pagination.NewPager(client, url, func(r pagination.PageResult) pagination.Page {
		return VolumePage{pagination.SinglePageBase(r)}
	})
}

func Detail(client *golangsdk.ServiceClient, opts ListOptsBuilder) pagination.Pager {
	url := detailURL(client)
	if opts != nil {
		query, err := opts.ToVolumeListQuery()
		if err != nil {
			return pagination.Pager{Err: err}
		}
		url += query
	}

	return pagination.NewPager(client, url, func(r pagination.PageResult) pagination.Page {
		return VolumePage{pagination.SinglePageBase(r)}
	})
}

// UpdateOptsBuilder allows extensions to add additional parameters to the
// Update request.
type UpdateOptsBuilder interface {
	ToVolumeUpdateMap() (map[string]interface{}, error)
}

// UpdateOpts contain options for updating an existing Volume. This object is passed
// to the volumes.Update function. For more information about the parameters, see
// the Volume object.
type UpdateOpts struct {
	Name        string            `json:"name,omitempty"`
	Description string            `json:"description,omitempty"`
	Metadata    map[string]string `json:"metadata,omitempty"`
}

// ToVolumeUpdateMap assembles a request body based on the contents of an
// UpdateOpts.
func (opts UpdateOpts) ToVolumeUpdateMap() (map[string]interface{}, error) {
	return golangsdk.BuildRequestBody(opts, "volume")
}

// Update will update the Volume with provided information. To extract the updated
// Volume from the response, call the Extract method on the UpdateResult.
func Update(client *golangsdk.ServiceClient, id string, opts UpdateOptsBuilder) (r UpdateResult) {
	b, err := opts.ToVolumeUpdateMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = client.Put(updateURL(client, id), b, &r.Body, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return
}

// IDFromName is a convienience function that returns a server's ID given its name.
func IDFromName(client *golangsdk.ServiceClient, name string) (string, error) {
	count := 0
	id := ""
	pages, err := List(client, nil).AllPages()
	if err != nil {
		return "", err
	}

	all, err := ExtractVolumes(pages)
	if err != nil {
		return "", err
	}

	for _, s := range all {
		if s.Name == name {
			count++
			id = s.ID
		}
	}

	switch count {
	case 0:
		return "", golangsdk.ErrResourceNotFound{Name: name, ResourceType: "volume"}
	case 1:
		return id, nil
	default:
		return "", golangsdk.ErrMultipleResourcesFound{
			Name:         name,
			Count:        count,
			ResourceType: "volume",
		}
	}
}

// MetadataOptsBuilder allows extensions to add additional parameters to
// the meatadata requests.
type MetadataOptsBuilder interface {
	ToVolumeMetadataMap() (map[string]interface{}, error)
}

// MetadataOpts contain options for creating or updating an existing Voulme. This
// object is passed to the volumes create and update function. For more information
// about the parameters, see the Volume object.
type MetadataOpts struct {
	Metadata map[string]string `json:"meta,omitempty"`
}

// ToSnapshotMetadataMap assembles a request body based on the contents of
// an MetadataOpts.
func (opts MetadataOpts) ToVolumeMetadataMap() (map[string]interface{}, error) {
	return golangsdk.BuildRequestBody(opts, "")
}

// CreateMetadata create metadata for Volume.
func CreateMetadata(client *golangsdk.ServiceClient, id string, opts MetadataOptsBuilder) (r MetadataResult) {
	b, err := opts.ToVolumeMetadataMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = client.Post(metadataURL(client, id), b, &r.Body, &golangsdk.RequestOpts{
		OkCodes: []int{202},
	})
	return
}

// GetMetadata returns exist metadata of Volume.
func GetMetadata(client *golangsdk.ServiceClient, id string) (r MetadataResult) {
	_, r.Err = client.Get(metadataURL(client, id), &r.Body, nil)
	return
}

// UpdateMetadata will update metadata according to request map.
func UpdateMetadata(client *golangsdk.ServiceClient, id string, opts MetadataOptsBuilder) (r MetadataResult) {
	b, err := opts.ToVolumeMetadataMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = client.Put(metadataURL(client, id), b, &r.Body, &golangsdk.RequestOpts{
		OkCodes: []int{202},
	})
	return
}

// GetMetadataKey return specific key value in metadata.
func GetMetadataKey(client *golangsdk.ServiceClient, id, key string) (r MetadataResult) {
	_, r.Err = client.Get(metadataKeyURL(client, id, key), &r.Body, nil)
	return
}

// UpdateMetadataKey update sepcific key to the given map key value.
func UpdateMetadataKey(client *golangsdk.ServiceClient, id, key string, opts MetadataOptsBuilder) (r MetadataResult) {
	b, err := opts.ToVolumeMetadataMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = client.Put(metadataKeyURL(client, id, key), b, &r.Body, &golangsdk.RequestOpts{
		OkCodes: []int{202},
	})
	return
}

// DeleteMetadataKey delete specific key in metadata
func DeleteMetadataKey(client *golangsdk.ServiceClient, id, key string) (r DeleteMetadataKeyResult) {
	_, r.Err = client.Delete(metadataKeyURL(client, id, key), nil)
	return
}

// ExtendSizeOptsBuilder allows extensions to add additional parameters to the
// ExtendSize request.
type ExtendSizeOptsBuilder interface {
	ToVolumeExtendSizeMap() (map[string]interface{}, error)
}

// ExtendSizeOpts contains options for extending the size of an existing Volume.
// This object is passed to the volumes.ExtendSize function.
type ExtendSizeOpts struct {
	// NewSize is the new size of the volume, in GB.
	NewSize int `json:"new_size" required:"true"`
}

// ToVolumeExtendSizeMap assembles a request body based on the contents of an
// ExtendSizeOpts.
func (opts ExtendSizeOpts) ToVolumeExtendSizeMap() (map[string]interface{}, error) {
	return golangsdk.BuildRequestBody(opts, "os-extend")
}

// ExtendSize will extend the size of the volume based on the provided information.
// This operation does not return a response body.
func ExtendSize(client *golangsdk.ServiceClient, id string, opts ExtendSizeOptsBuilder) (r ExtendSizeResult) {
	b, err := opts.ToVolumeExtendSizeMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = client.Post(actionURL(client, id), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{202},
	})
	return
}

// SetBootableOptsBuilder allows extensions to add additional parameters to the
// SetBootable request.
type SetBootableOptsBuilder interface {
	ToVolumeSetBootableMap() (map[string]interface{}, error)
}

// SetBootableOpts contains options for setting bootable flag of an existing Volume.
// This object is passed to the volumes.SetBootable function.
type SetBootableOpts struct {
	// Bootable is bool of true or false
	Bootable bool `json:"bootable" required:"true"`
}

// ToVolumeSetBootableMap assembles a request body based on the contents of an
// SetBootableOpts.
func (opts SetBootableOpts) ToVolumeSetBootableMap() (map[string]interface{}, error) {
	return golangsdk.BuildRequestBody(opts, "os-set_bootable")
}

// SetBootable will set bootable flag of the volume based on the provided information.
// This operation does not return a response body.
func SetBootable(client *golangsdk.ServiceClient, id string, opts SetBootableOptsBuilder) (r SetBootableResult) {
	b, err := opts.ToVolumeSetBootableMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = client.Post(actionURL(client, id), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{202},
	})
	return
}

// SetReadOnlyOptsBuilder allows extensions to add additional parameters to the
// SetReadOnly request.
type SetReadOnlyOptsBuilder interface {
	ToVolumeSetReadOnlyMap() (map[string]interface{}, error)
}

// SetReadOnlyOpts contains options for setting readonly flag of an existing Volume.
// This object is passed to the volumes.SetReadOnly function.
type SetReadOnlyOpts struct {
	// ReadOnly is bool of true or false
	ReadOnly bool `json:"readonly" required:"true"`
}

// ToVolumeSetReadOnlyMap assembles a request body based on the contents of an
// SetReadOnlyOpts.
func (opts SetReadOnlyOpts) ToVolumeSetReadOnlyMap() (map[string]interface{}, error) {
	return golangsdk.BuildRequestBody(opts, "os-update_readonly_flag")
}

// SetReadOnly will set readonly flag of the volume based on the provided information.
// This operation does not return a response body.
func SetReadOnly(client *golangsdk.ServiceClient, id string, opts SetReadOnlyOptsBuilder) (r SetReadOnlyResult) {
	b, err := opts.ToVolumeSetReadOnlyMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = client.Post(actionURL(client, id), b, nil, &golangsdk.RequestOpts{
		OkCodes: []int{202},
	})
	return
}
