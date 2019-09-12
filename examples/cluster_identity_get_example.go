package main

import (
	"fmt"
	"time"
	"go-ontap-sdk/ontap"
)

func main() {
	c := ontap.NewClient(
	    "https://mycluster.example.com",
	    &ontap.ClientOptions {
		Version: "1.160",
		BasicAuthUser: "admin",
		BasicAuthPassword: "secret",
		SSLVerify: false,
		Debug: false,
    		Timeout: 60 * time.Second,
	    },
	)
	options := &ontap.ClusterIdentityOptions {}
	response, _, err := c.ClusterIdentityGetAPI(options)
	if err != nil {
		fmt.Print(err)
	} else {
		fmt.Printf("%s\n", response.Results.ClusterIdentityInfo[0].ClusterName)
		fmt.Printf("\tLocation: %s\n", response.Results.ClusterIdentityInfo[0].ClusterLocation)
	}
}
