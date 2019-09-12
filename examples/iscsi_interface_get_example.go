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
	options := &ontap.IscsiInterfaceGetOptions {
			MaxRecords: 1024,
			Query: &ontap.IscsiInterfaceQuery {
				IscsiInterfaceListEntryInfo: &ontap.IscsiInterfaceListEntryInfo {
				},
			},
	}
	//c.SetVserver("myvserver")
	response, _, err := c.IscsiInterfaceGetAPI(options)
	if err != nil {
		fmt.Print(err)
	} else {
		if response.Results.NumRecords > 0 {
    			for _, iface := range response.Results.AttributesList.IscsiInterfaceAttributes {
				fmt.Printf("%s\n", iface.InterfaceName)
				fmt.Printf("\tCurrent Node: %s\n", iface.CurrentNode)
				fmt.Printf("\tCurrent Port: %s\n", iface.CurrentPort)
				fmt.Printf("\tIP Address: %s\n", iface.IpAddress)
				fmt.Printf("\tIP Port: %s\n", iface.IpPort)
				fmt.Printf("\tTpgroup Name: %s\n", iface.TpgroupName)
			}
		}
		fmt.Printf("Total Records: %d\n", response.Results.NumRecords)
	}
}
