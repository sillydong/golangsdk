/*
Package quotasets enables retrieving and managing Block Storage quotas.

Example to Get Quota Set Usage

	quotaset, err := quotasets.GetUsage(blockStorageClient, "project-id").Extract()
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", quotaset)

*/
package quotasets
