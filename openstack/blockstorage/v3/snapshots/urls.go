package snapshots

import "github.com/bizflycloud/gophercloud"

func createURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("snapshots")
}

func deleteURL(c *gophercloud.ServiceClient, id string) string {
	return c.ServiceURL("snapshots", id)
}

func getURL(c *gophercloud.ServiceClient, id string) string {
	return deleteURL(c, id)
}

func listURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("snapshots", "detail")
}

func updateURL(c *gophercloud.ServiceClient, id string) string {
	return deleteURL(c, id)
}

func metadataURL(c *gophercloud.ServiceClient, id string) string {
	return c.ServiceURL("snapshots", id, "metadata")
}

func updateMetadataURL(c *gophercloud.ServiceClient, id string) string {
	return metadataURL(c, id)
}

func actionURL(c *gophercloud.ServiceClient, id string) string {
	return c.ServiceURL("snapshots", id, "action")
}

func metadatumURL(c *gophercloud.ServiceClient, id, key string) string {
	return c.ServiceURL("snapshots", id, "metadata", key)
}
