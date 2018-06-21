/*
Package cloudimages enables management and retrieval of images from the
OpenStack Image Service.

Example to List Image Tags

	r := cloudimages.ListImageTags(client, ListImageTagsOpts{})

	list, err := r.Extract()
	if err != nil {
		panic(err)
	}

Examples to Set Image Tag

	imageID := "imageID"
	tag := "imageTag"
	r := cloudimages.SetImageTag(client, SetImageTagOpts{
		ImageID: imageID,
		Tag: tag,
	})
	if r.Err != nil {
		panic(r.Err)
	}

Example to Import Image

	imageID := "imageID"
	imageURL := "imageURL"
	r := cloudimages.ImportImage(client, imageID, imageURL)
	if r.Err != nil {
		panic(r.Err)
	}

Example to Export Image

	imageID := "imageID"
	bucketURL := "bucketURL"
	fileFormat := "fileFormat"
	r := cloudimages.ExportImage(client, imageID, bucketURL, fileFormat)
	if r.Err != nil {
		panic(r.Err)
	}

Example to Copy Image

	imageID := "imageID"
	r := cloudimages.CopyImage(client, imageID, CopyImageOpts{
		Name: "image_name",
	})
	if r.Err != nil {
		panic(err)
	}

Example to add image members

	r := cloudimages.AddImageMembers(client, AddImageMembersOpts{
		Images: []string{"image1"},
		Projects: []string{"project1"},
	})
	if r.Err != nil {
		panic(r.Err)
	}

Example to update image members

	r := cloudimages.UpdateImageMembers(client, UpdateImageMembersOpts{
		Images: []string{"image1"},
		Projects: "project1",
		Status: "accepted",
	})
	if r.Err != nil {
		panic(r.Err)
	}

Example to delete image members

	r := cloudimages.DeleteImageMembers(client, DeleteImageMembersOpts{
		Images: []string{"image1"},
		Projects: []string{"project1"},
	})
	if r.Err != nil {
		panic(r.Err)
	}
*/
package cloudimages
