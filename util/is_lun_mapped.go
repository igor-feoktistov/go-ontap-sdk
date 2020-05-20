package util

import (
	"go-ontap-sdk/ontap"
)

func IsLunMapped(c *ontap.Client, lunPath string, igroupName string) (mapped bool, err error) {
	response, _, err := c.LunMapListInfoAPI(lunPath)
	if err == nil {
		for _, igroup := range response.Results.InitiatorGroups.IgroupAttributes {
			if igroupName == igroup.InitiatorGroupName {
				mapped = true
			}
		}
	}
	return
}
