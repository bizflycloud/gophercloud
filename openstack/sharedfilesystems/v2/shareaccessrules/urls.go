package shareaccessrules

import (
	"github.com/bizflycloud/gophercloud"
)

const shareAccessRulesEndpoint = "share-access-rules"

func getURL(c *gophercloud.ServiceClient, accessID string) string {
	return c.ServiceURL(shareAccessRulesEndpoint, accessID)
}