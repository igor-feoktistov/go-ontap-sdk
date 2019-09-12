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
	options := &ontap.IscsiConnectionGetOptions {
			MaxRecords: 1024,
			Query: &ontap.IscsiConnectionQuery {
				IscsiConnectionListEntryInfo: &ontap.IscsiConnectionListEntryInfo {
					RemoteIpAddress: "192.168.1.5",
				},
			},
	}
	//c.SetVserver("myvserver")
	response, _, err := c.IscsiConnectionGetAPI(options)
	if err != nil {
		fmt.Print(err)
	} else {
		if response.Results.NumRecords > 0 {
    			for _, conn := range response.Results.AttributesList.IscsiConnectionAttributes {
				fmt.Printf("Connection ID %d\n", conn.ConnectionId)
				fmt.Printf("\tConnection State: %s\n", conn.ConnectionState)
				fmt.Printf("\tRemore IP: %s\n", conn.RemoteIpAddress)
				fmt.Printf("\tRemore IP Port: %d\n", conn.RemoteIpPort)
				fmt.Printf("\tLIF: %s\n", conn.InterfaceName)
				fmt.Printf("\tTpgroup Name: %s\n", conn.TpgroupName)
			}
		}
		fmt.Printf("Total Records: %d\n", response.Results.NumRecords)
	}
}
