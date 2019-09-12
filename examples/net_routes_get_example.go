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
	options := &ontap.NetRoutesGetOptions {
			MaxRecords: 1024,
			Query: &ontap.NetRoutesQuery {
				NetVsRoutesInfo: &ontap.NetVsRoutesInfo {
				},
			},
	}
	//c.SetVserver("myvserver")
	response, _, err := c.NetRoutesGetAPI(options)
	if err != nil {
		fmt.Print(err)
	} else {
		if response.Results.NumRecords > 0 {
    			for _, route := range response.Results.AttributesList.NetRoutesAttributes {
				fmt.Printf("%s\n", route.Destination)
				fmt.Printf("\tGateway: %s\n", route.Gateway)
				fmt.Printf("\tMetric: %d\n", route.Metric)
			}
		}
		fmt.Printf("Total Records: %d\n", response.Results.NumRecords)
	}
}
