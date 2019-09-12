package main

import (
	"fmt"
	"time"
	"go-ontap-sdk/ontap"
)

func main() {
	var numRecords int
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
			MaxRecords: 512,
			Query: &ontap.VolumeQuery {
				VolumeInfo: &ontap.VolumeInfo {},
			},
	}
	//c.SetVserver("myverver")
	response, err := c.VolumeGetIterAPI(options)
	if err != nil {
	    fmt.Print(err.Error())
	} else {
		for _, responseVolume := range response {
			numRecords += responseVolume.Results.NumRecords
			for _, volume := range responseVolume.Results.AttributesList {
				fmt.Printf("%s\n", volume.VolumeIDAttributes.Name)
				fmt.Printf("\tNode: %s\n", volume.VolumeIDAttributes.Node)
				fmt.Printf("\tAggregate: %s\n", volume.VolumeIDAttributes.ContainingAggregateName)
				fmt.Printf("\tSize: %.2fGB\n", float64(volume.VolumeSpaceAttributes.Size)/1024/1024/1024)
			}
		}
		fmt.Printf("Total Records: %d\n", numRecords)
	}
}
