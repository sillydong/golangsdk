package apiversions

import (
	"net/url"
	"strings"

	"github.com/huaweicloud/golangsdk"
)

// getURL generates URL fot getting detail of specified version
func getURL(c *golangsdk.ServiceClient, version string) string {
	u, _ := url.Parse(c.ServiceURL(""))
	u.Path = "/" + strings.TrimRight(version, "/") + "/"
	return u.String()
}

// listURL generates URL for list versions
func listURL(c *golangsdk.ServiceClient) string {
	u, _ := url.Parse(c.ServiceURL(""))
	u.Path = "/"
	return u.String()
}
