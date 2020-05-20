package util

import (
	"fmt"
	"time"
	"github.com/igor-feoktistov/go-ontap-sdk/ontap"
)

func LunCopy(c *ontap.Client, srcLunPath string, dstLunPath string) (err error) {
	var lunCopyInfo ontap.LunCopyInfo
	optionsStart := &ontap.LunCopyStartOptions {
			    Paths: &[]ontap.LunPathPair {
					ontap.LunPathPair {
						DestinationPath: dstLunPath,
						SourcePath: srcLunPath,
					},
			    },
	}
	responseStart, _, err := c.LunCopyStartAPI(optionsStart)
	if err != nil {
		return
	} else {
		optionsGet := &ontap.LunCopyGetOptions {
				MaxRecords: 1,
				Query: &ontap.LunCopyInfo {
					JobUuid: responseStart.Results.JobUuid,
				},
		}
		for {
			responseGet, _, err := c.LunCopyGetAPI(optionsGet)
			if err != nil {
				break
			} else {
				if responseGet.Results.NumRecords > 0 {
    					lunCopyInfo = responseGet.Results.AttributesList.LunCopyAttributes[0]
    					if lunCopyInfo.JobStatus == "complete" || lunCopyInfo.JobStatus == "paused_admin" || lunCopyInfo.JobStatus == "paused_error" || lunCopyInfo.JobStatus == "destroyed" {
						break
					} else {
						time.Sleep(time.Second)
					}
				} else {
					err = fmt.Errorf("LunCopy: no LunCopyInfo found for Job UUID %s", responseStart.Results.JobUuid)
					break
				}
			}
		}
	}
	if err == nil {
		if lunCopyInfo.JobStatus != "complete" {
			err = fmt.Errorf("LunCopy: completed with failure: %s", lunCopyInfo.LastFailureReason)
		}
	}
	return
}
