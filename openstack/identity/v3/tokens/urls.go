package tokens

import "github.com/huaweicloud/golangsdk"

// tokenURL generate url to obtaining or get the user token
func tokenURL(c *golangsdk.ServiceClient) string {
	return c.ServiceURL("auth", "tokens")
}
