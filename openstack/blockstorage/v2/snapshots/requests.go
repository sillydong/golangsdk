package snapshots

import (
	"github.com/huaweicloud/golangsdk"
	"github.com/huaweicloud/golangsdk/pagination"
)

// CreateOptsBuilder allows extensions to add additional parameters to the
// Create request.
type CreateOptsBuilder interface {
	ToSnapshotCreateMap() (map[string]interface{}, error)
}

// CreateOpts contains options for creating a Snapshot. This object is passed to
// the snapshots.Create function. For more information about these parameters,
// see the Snapshot object.
type CreateOpts struct {
	VolumeID    string            `json:"volume_id" required:"true"`
	Force       bool              `json:"force,omitempty"`
	Name        string            `json:"name,omitempty"`
	Description string            `json:"description,omitempty"`
	Metadata    map[string]string `json:"metadata"`
}

// ToSnapshotCreateMap assembles a request body based on the contents of a
// CreateOpts.
func (opts CreateOpts) ToSnapshotCreateMap() (map[string]interface{}, error) {
	return golangsdk.BuildRequestBody(opts, "snapshot")
}

// Create will create a new Snapshot based on the values in CreateOpts. To
// extract the Snapshot object from the response, call the Extract method on the
// CreateResult.
func Create(client *golangsdk.ServiceClient, opts CreateOptsBuilder) (r CreateResult) {
	b, err := opts.ToSnapshotCreateMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = client.Post(createURL(client), b, &r.Body, &golangsdk.RequestOpts{
		OkCodes: []int{202},
	})
	return
}

// Delete will delete the existing Snapshot with the provided ID.
func Delete(client *golangsdk.ServiceClient, id string) (r DeleteResult) {
	_, r.Err = client.Delete(deleteURL(client, id), nil)
	return
}

// Get retrieves the Snapshot with the provided ID. To extract the Snapshot
// object from the response, call the Extract method on the GetResult.
func Get(client *golangsdk.ServiceClient, id string) (r GetResult) {
	_, r.Err = client.Get(getURL(client, id), &r.Body, nil)
	return
}

// ListOptsBuilder allows extensions to add additional parameters to the List
// request.
type ListOptsBuilder interface {
	ToSnapshotListQuery() (string, error)
}

// ListOpts hold options for listing Snapshots. It is passed to the
// snapshots.List function.
type ListOpts struct {
	// Name will filter by the specified snapshot name.
	Name string `q:"name"`

	// Status will filter by the specified status.
	Status string `q:"status"`

	// SnapshotID will filter by a specified volume ID.
	SnapshotID string `q:"volume_id"`

	// Limit of pagination
	Limit int `q:"limit"`

	// Offset of pagination
	Offset int `q:"offset"`
}

// ToSnapshotListQuery formats a ListOpts into a query string.
func (opts ListOpts) ToSnapshotListQuery() (string, error) {
	q, err := golangsdk.BuildQueryString(opts)
	return q.String(), err
}

// List returns Snapshots optionally limited by the conditions provided in
// ListOpts.
func List(client *golangsdk.ServiceClient, opts ListOptsBuilder) pagination.Pager {
	url := listURL(client)
	if opts != nil {
		query, err := opts.ToSnapshotListQuery()
		if err != nil {
			return pagination.Pager{Err: err}
		}
		url += query
	}
	return pagination.NewPager(client, url, func(r pagination.PageResult) pagination.Page {
		return SnapshotPage{pagination.SinglePageBase(r)}
	})
}

func Detail(client *golangsdk.ServiceClient, opts ListOptsBuilder) pagination.Pager {
	url := detailURL(client)
	if opts != nil {
		query, err := opts.ToSnapshotListQuery()
		if err != nil {
			return pagination.Pager{Err: err}
		}
		url += query
	}

	return pagination.NewPager(client, url, func(r pagination.PageResult) pagination.Page {
		return SnapshotPage{pagination.SinglePageBase(r)}
	})
}

type UpdateOptsBuilder interface {
	ToSnapshotUpdateMap() (map[string]interface{}, error)
}

type UpdateOpts struct {
	Name               string `json:"name,omitempty"`
	Description        string `json:"description,omitempty"`
	DisplayName        string `json:"display_name,omitempty"`
	DisplayDescription string `json:"display_description,omitempty"`
}

func (opts UpdateOpts) ToSnapshotUpdateMap() (map[string]interface{}, error) {
	return golangsdk.BuildRequestBody(opts, "snapshot")
}

func Update(client *golangsdk.ServiceClient, id string, opts UpdateOptsBuilder) (r UpdateResult) {
	b, err := opts.ToSnapshotUpdateMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = client.Put(updateURL(client, id), b, &r.Body, &golangsdk.RequestOpts{
		OkCodes: []int{200},
	})
	return
}

type RollbackOptsBuilder interface {
	ToSnapshotRollbackMap() (map[string]interface{}, error)
}

type RollbackOpts struct {
	VolumeID string `json:"volume_id,omitempty"`
	Name     string `json:"name,omitempty"`
}

func (opts RollbackOpts) ToSnapshotRollbackMap() (map[string]interface{}, error) {
	return golangsdk.BuildRequestBody(opts, "rollback")
}

func Rollback(client *golangsdk.ServiceClient, id string, opts RollbackOptsBuilder) (r RollbackResult) {
	b, err := opts.ToSnapshotRollbackMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = client.Post(rollbackURL(client, id), b, &r.Body, &golangsdk.RequestOpts{
		OkCodes: []int{202},
	})
	return
}

// MetadataOptsBuilder allows extensions to add additional parameters to
// the meatadata requests.
type MetadataOptsBuilder interface {
	ToSnapshotMetadataMap() (map[string]interface{}, error)
}

// MetadataOpts contain options for creating or updating an existing Voulme. This
// object is passed to the volumes create and update function. For more information
// about the parameters, see the Snapshot object.
type MetadataOpts struct {
	Metadata map[string]string `json:"metadata,omitempty"`
}

// ToSnapshotMetadataMap assembles a request body based on the contents of
// an MetadataOpts.
func (opts MetadataOpts) ToSnapshotMetadataMap() (map[string]interface{}, error) {
	return golangsdk.BuildRequestBody(opts, "")
}

// CreateMetadata create metadata for Snapshot.
func CreateMetadata(client *golangsdk.ServiceClient, id string, opts MetadataOptsBuilder) (r MetadataResult) {
	b, err := opts.ToSnapshotMetadataMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = client.Post(metadataURL(client, id), b, &r.Body, &golangsdk.RequestOpts{
		OkCodes: []int{202},
	})
	return
}

// GetMetadata returns exist metadata of Snapshot.
func GetMetadata(client *golangsdk.ServiceClient, id string) (r MetadataResult) {
	_, r.Err = client.Get(metadataURL(client, id), &r.Body, nil)
	return
}

// UpdateMetadata will update metadata according to request map.
func UpdateMetadata(client *golangsdk.ServiceClient, id string, opts MetadataOptsBuilder) (r MetadataResult) {
	b, err := opts.ToSnapshotMetadataMap()
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
	b, err := opts.ToSnapshotMetadataMap()
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

// IDFromName is a convienience function that returns a snapshot's ID given its name.
func IDFromName(client *golangsdk.ServiceClient, name string) (string, error) {
	count := 0
	id := ""
	pages, err := List(client, nil).AllPages()
	if err != nil {
		return "", err
	}

	all, err := ExtractSnapshots(pages)
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
		return "", golangsdk.ErrResourceNotFound{Name: name, ResourceType: "snapshot"}
	case 1:
		return id, nil
	default:
		return "", golangsdk.ErrMultipleResourcesFound{
			Name:         name,
			Count:        count,
			ResourceType: "snapshot",
		}
	}
}
