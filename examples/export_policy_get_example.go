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
	options := &ontap.ExportPolicyGetOptions {
		    MaxRecords: 1024,
            	    Query: &ontap.ExportPolicyQuery {
                	    ExportPolicyInfo: &ontap.ExportPolicyInfo {
				    Vserver: "svmdurlabgen01spd",
                            },
                    },
	}
	//c.SetVserver("myvserver")
	response, _, err := c.ExportPolicyGetAPI(options)
	if err != nil {
		fmt.Print(err)
	} else {
		if response.Results.NumRecords > 0 {
    			for _, policy := range response.Results.AttributesList.ExportPolicyAttributes {
				fmt.Printf("%s\n", policy.PolicyName)
			}
		}
		fmt.Printf("Total Records: %d\n", response.Results.NumRecords)
	}
}
