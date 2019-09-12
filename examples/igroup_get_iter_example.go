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
	options := &ontap.IgroupGetOptions {
			MaxRecords: 16,
	}
	//c.SetVserver("myvserver")
	response, err := c.IgroupGetIterAPI(options)
	if err != nil {
	    fmt.Print(err.Error())
	} else {
		numRecords = 0
		for _, responseIgroup := range response {
			numRecords += responseIgroup.Results.NumRecords
    			for _, igroup := range responseIgroup.Results.AttributesList.IgroupAttributes {
				fmt.Printf("%s\n", igroup.InitiatorGroupName)
				fmt.Printf("\tOS Type: %s\n", igroup.InitiatorGroupOsType)
				fmt.Println("\tInitators:")
    				for _, initiator := range igroup.Initiators {
					fmt.Printf("\t\tInitiator Name: %s\n", initiator.InitiatorName)
    				}
			}
		}
		fmt.Printf("Total Records: %d\n", numRecords)
	}
}
