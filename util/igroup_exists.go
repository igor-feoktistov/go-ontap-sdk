package util

import (
	"go-ontap-sdk/ontap"
)

func IgroupExists(c *ontap.Client, igroupName string) (exists bool, err error) {
	options := &ontap.IgroupGetOptions {
			MaxRecords: 1024,
			Query: &ontap.IgroupQuery {
				IgroupInfo: &ontap.IgroupInfo {
					InitiatorGroupName: igroupName,
				},
			},
	}
	response, _, err := c.IgroupGetAPI(options)
	if err == nil {
		if response.Results.NumRecords > 0 {
			exists = true
		}
	}
	return
}
