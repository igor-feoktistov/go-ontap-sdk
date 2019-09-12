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
	options := &ontap.IscsiInitiatorGetOptions {
			MaxRecords: 1024,
			Query: &ontap.IscsiInitiatorQuery {
				IscsiInitiatorListEntryInfo: &ontap.IscsiInitiatorListEntryInfo {
					InitiatorNodename: "iqn.1998-01.com.vmware:myhost01-dev",
				},
			},
	}
	//c.SetVserver("vsadmin")
	response, _, err := c.IscsiInitiatorGetAPI(options)
	if err != nil {
		fmt.Print(err)
	} else {
		if response.Results.NumRecords > 0 {
    			for _, initiator := range response.Results.AttributesList.IscsiInitiatorAttributes {
				fmt.Printf("%s\n", initiator.InitiatorNodename)
				fmt.Printf("\tISID: %s\n", initiator.Isid)
				fmt.Printf("\tTpgroup Name: %s\n", initiator.TpgroupName)
			}
		}
		fmt.Printf("Total Records: %d\n", response.Results.NumRecords)
        }
}
