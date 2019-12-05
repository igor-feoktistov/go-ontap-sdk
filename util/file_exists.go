package util

import (
	"go-ontap-sdk/ontap"
)

func FileExists(c *ontap.Client, filePath string) (exists bool, err error) {
	response, _, err := c.FileGetFileInfoAPI(filePath)
	if err == nil {
		exists = true
	} else {
		if response.Results.ErrorNo == ontap.EONTAPI_ENOENT {
			err = nil
		}
	}
	return
}
