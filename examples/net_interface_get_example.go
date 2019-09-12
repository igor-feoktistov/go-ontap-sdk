package main

import (
	"fmt"
	"time"
	"net"
	"strings"
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
	options := &ontap.NetInterfaceGetOptions {
			MaxRecords: 1024,
			Query: &ontap.NetInterfaceQuery {
				NetInterfaceInfo: &ontap.NetInterfaceInfo {
					DataProtocols: &[]string{"nfs"},
				},
			},
	}
	//c.SetVserver("myvserver")
	response, _, err := c.NetInterfaceGetAPI(options)
	if err != nil {
		fmt.Print(err)
	} else {
		if response.Results.NumRecords > 0 {
    			for _, lif := range response.Results.AttributesList.NetInterfaceAttributes {
				fmt.Printf("%s\n", lif.InterfaceName)
				fmt.Printf("\tAddress: %s\n", lif.Address)
				fmt.Printf("\tNetmask: %s\n", lif.Netmask)
				fmt.Printf("\tNetwork: %s/%d\n", net.ParseIP(lif.Address).Mask(net.CIDRMask(lif.NetmaskLength, 32)), lif.NetmaskLength)
				fmt.Printf("\tProtocols: %s\n", strings.Join(*lif.DataProtocols, ","))
				fmt.Printf("\tHome Node: %s\n", lif.HomeNode)
				fmt.Printf("\tHome Port: %s\n", lif.HomePort)
			}
		}
		fmt.Printf("Total Records: %d\n", response.Results.NumRecords)
	}
}
