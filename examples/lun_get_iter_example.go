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
		BasicAuthUser: "vsadmin",
		BasicAuthPassword: "secret",
		SSLVerify: false,
		Debug: false,
    		Timeout: 60 * time.Second,
		Version: "1.160",
	    },
	)
	options := &ontap.LunGetOptions {
			MaxRecords: 256,
	}
	//c.SetVserver("myvserver")
	response, err := c.LunGetIterAPI(options)
	if err != nil {
		fmt.Print(err.Error())
	} else {
		numRecords = 0
		for _, responseLun := range response {
			numRecords += responseLun.Results.NumRecords
    			for _, lun := range responseLun.Results.AttributesList.LunAttributes {
				fmt.Printf("%s\n", lun.Path)
				fmt.Printf("\tNode: %s\n", lun.Node)
				fmt.Printf("\tVolume: %s\n", lun.Volume)
				fmt.Printf("\tSize: %.2fGB\n", float64(lun.Size)/1024/1024/1024)
			}
		}
		fmt.Printf("Total Records: %d\n", numRecords)
	}
}
