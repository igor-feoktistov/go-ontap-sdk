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
	options := &ontap.ExportPolicyGetOptions {
			MaxRecords: 1024,
            		Query: &ontap.ExportPolicyQuery {
                		ExportPolicyInfo: &ontap.ExportPolicyInfo {
					Vserver: "myvserver",
                        	},
                	},
	}
	//c.SetVserver("myvserver")
	response, err := c.ExportPolicyGetIterAPI(options)
	if err != nil {
		fmt.Print(err)
	} else {
		numRecords = 0
        	for _, responsePolicy := range response {
                	numRecords += responsePolicy.Results.NumRecords
    			for _, policy := range responsePolicy.Results.AttributesList.ExportPolicyAttributes {
				fmt.Printf("%s\n", policy.PolicyName)
			}
		}
		fmt.Printf("Total Records: %d\n", numRecords)
	}
}
