package images

import (
	"net/url"

	"github.com/huaweicloud/golangsdk"
)

// `listURL` is a pure function. `listURL(c)` is a URL for which a GET
// request will respond with a list of images in the service `c`.
func listURL(c *golangsdk.ServiceClient) string {
	return c.ServiceURL("images")
}

// `listCloudImagesURL` is a pure function. `listCloudImagesURL(c)` is a URL for
// which a GET request will respond with a list of images in the service `c`.
func listCloudImagesURL(c *golangsdk.ServiceClient) string {
	return c.ServiceURL("cloudimages")
}

// createURL generates a url for creating the metadata of an image
func createURL(c *golangsdk.ServiceClient) string {
	return c.ServiceURL("images")
}

// `imageURL(c,i)` is the URL for the image identified by ID `i` in
// the service `c`.
func imageURL(c *golangsdk.ServiceClient, imageID string) string {
	return c.ServiceURL("images", imageID)
}

// `getURL(c,i)` is a URL for which a GET request will respond with
// information about the image identified by ID `i` in the service
// `c`.
func getURL(c *golangsdk.ServiceClient, imageID string) string {
	return imageURL(c, imageID)
}

// updateURL update the attributes of image
func updateURL(c *golangsdk.ServiceClient, imageID string) string {
	return imageURL(c, imageID)
}

// `updateCloudImageURL` is a function to generate a url for update image
// attributes
func updateCloudImageURL(c *golangsdk.ServiceClient, imageID string) string {
	return c.ServiceURL("cloudimages", imageID)
}

// actionCloudImageURL generate url to do the action of cloudimages
func actionCloudImageURL(c *golangsdk.ServiceClient) string {
	return c.ServiceURL("cloudimages", "action")
}

// `deleteURL` is a function to generate a url for image deletion
func deleteURL(c *golangsdk.ServiceClient, imageID string) string {
	return imageURL(c, imageID)
}

// builds next page full url based on current url
func nextPageURL(currentURL string, next string) (string, error) {
	base, err := url.Parse(currentURL)
	if err != nil {
		return "", err
	}
	rel, err := url.Parse(next)
	if err != nil {
		return "", err
	}
	return base.ResolveReference(rel).String(), nil
}

// putTagURL add a tag to the image
func putTagURL(c *golangsdk.ServiceClient, imageID, tag string) string {
	return c.ServiceURL("images", imageID, "tags", tag)
}

// deleteTagURL remove the tag from a image
func deleteTagURL(c *golangsdk.ServiceClient, imageID, tag string) string {
	return c.ServiceURL("images", imageID, "tags", tag)
}

// getImageSchemas generate url to get image schemas
func getImageSchemas(c *golangsdk.ServiceClient) string {
	return c.ServiceURL("schemas", "image")
}

// getImagesSchemas generate url to get images schemas
func getImagesSchemas(c *golangsdk.ServiceClient) string {
	return c.ServiceURL("schemas", "images")
}
