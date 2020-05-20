package util

import (
	"github.com/igor-feoktistov/go-ontap-sdk/ontap"
)

func VolumeExists(c *ontap.Client, volumeName string) (exists bool, err error) {
	options := &ontap.VolumeGetOptions {
			MaxRecords: 1024,
			Query: &ontap.VolumeQuery {
				VolumeInfo: &ontap.VolumeInfo {
						VolumeIDAttributes: &ontap.VolumeIDAttributes {
								Name: volumeName,
						},
				},
			},
	}
	response, _, err := c.VolumeGetAPI(options)
	if err == nil {
		if response.Results.NumRecords > 0 {
			exists = true
		}
	}
	return
}
