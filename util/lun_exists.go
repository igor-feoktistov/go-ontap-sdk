package util

import (
	"github.com/igor-feoktistov/go-ontap-sdk/ontap"
)

func LunExists(c *ontap.Client, lunPath string) (exists bool, err error) {
	options := &ontap.LunGetOptions {
			MaxRecords: 1024,
			Query: &ontap.LunQuery {
				LunInfo: &ontap.LunInfo {
					Path: lunPath,
				},
			},
	}
	response, _, err := c.LunGetAPI(options)
	if err == nil {
		if response.Results.NumRecords > 0 {
			exists = true
		}
	}
	return
}
