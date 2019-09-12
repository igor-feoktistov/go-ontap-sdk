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
	options := &ontap.VolumeGetOptions {
			MaxRecords: 1024,
			Query: &ontap.VolumeQuery {
                		VolumeInfo: &ontap.VolumeInfo {
					VolumeIDAttributes: &ontap.VolumeIDAttributes {
                        			Name: "vol_go_test01",
                        		},
                    		},
			},
	}
	//c.SetVserver("myvserver")
	response, _, err := c.VolumeGetAPI(options)
	if err != nil {
		fmt.Print(err.Error())
	} else {
		if response.Results.NumRecords > 0 {
    			for _, volume := range response.Results.AttributesList {
				fmt.Printf("%s\n", volume.VolumeIDAttributes.Name)
				fmt.Printf("\tNode: %s\n", volume.VolumeIDAttributes.Node)
				fmt.Printf("\tAggregate: %s\n", volume.VolumeIDAttributes.ContainingAggregateName)
				fmt.Printf("\tSize: %.2fGB\n", float64(volume.VolumeSpaceAttributes.Size)/1024/1024/1024)
			}
			fmt.Printf("Total Records: %d\n", response.Results.NumRecords)
		}
	}
}
