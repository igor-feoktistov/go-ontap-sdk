package main

import (
	"fmt"
	"time"
	"go-ontap-sdk/ontap"
)

func main() {
	c := ontap.NewClient(
	    "https://myvserver.example.com",
	    &ontap.ClientOptions {
		Version: "1.160",
		BasicAuthUser: "vsadmin",
		BasicAuthPassword: "secret",
		SSLVerify: false,
		Debug: false,
    		Timeout: 60 * time.Second,
	    },
	)
	lunPathPair := ontap.LunPathPair {
			DestinationPath: "/vol/vol_go_test02/lun_go_test02",
			SourcePath: "/vol/vol_go_test01/lun_go_test01",
	}
	optionsStart := &ontap.LunCopyStartOptions {
			    Paths: &[]ontap.LunPathPair {lunPathPair},
	}
	//c.SetVserver("myvserver")
	responseStart, _, err := c.LunCopyStartAPI(optionsStart)
	if err != nil {
		fmt.Print(err.Error())
	} else {
		fmt.Printf("LUN copy started, Job UUID %s, waiting for completion", responseStart.Results.JobUuid)
		optionsGet := &ontap.LunCopyGetOptions {
				    MaxRecords: 1,
				    Query: &ontap.LunCopyInfo {
					    JobUuid: responseStart.Results.JobUuid,
				    },
		}
		for {
			fmt.Printf(".")
			responseGet, _, err := c.LunCopyGetAPI(optionsGet)
			if err != nil {
				fmt.Print(err)
				break
			} else {
				if responseGet.Results.NumRecords > 0 {
    					copyInfo := responseGet.Results.AttributesList.LunCopyAttributes[0]
    					if copyInfo.JobStatus == "complete" {
						fmt.Printf("\nCompleted: Job UUID %s\n", copyInfo.JobUuid)
						fmt.Printf("\tStatus: %s\n", copyInfo.JobStatus)
						fmt.Printf("\tSource Path: %s\n", copyInfo.SourcePath)
						fmt.Printf("\tDestination Path: %s\n", copyInfo.DestinationPath)
						fmt.Printf("\tCopied Bytes: %d\n", copyInfo.ScannerTotal)
						break
					} else {
						time.Sleep(time.Second)
					}
				} else {
					break
				}
			}
		}
	}
}
