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
		BasicAuthUser: "vsadmin",
		BasicAuthPassword: "secret",
		SSLVerify: false,
		Debug: false,
    		Timeout: 60 * time.Second,
		Version: "1.160",
	    },
	)
	options := &ontap.IgroupGetOptions {
			MaxRecords: 1024,
			Query: &ontap.IgroupQuery {
				IgroupInfo: &ontap.IgroupInfo {
					InitiatorGroupName: "myhost01_dev_igroup01",
				},
			},
	}
	//c.SetVserver("myvserver")
	response, _, err := c.IgroupGetAPI(options)
	if err != nil {
		fmt.Print(err)
	} else {
		if response.Results.NumRecords > 0 {
    			for _, igroup := range response.Results.AttributesList.IgroupAttributes {
				fmt.Printf("%s\n", igroup.InitiatorGroupName)
				fmt.Printf("\tOS Type: %s\n", igroup.InitiatorGroupOsType)
				fmt.Println("\tInitators:")
    				for _, initiator := range igroup.Initiators {
					fmt.Printf("\t\tInitiator Name: %s\n", initiator.InitiatorName)
    				}
			}
		}
		fmt.Printf("Total Records: %d\n", response.Results.NumRecords)
	}
}
