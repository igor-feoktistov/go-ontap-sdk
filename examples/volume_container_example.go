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
		Version: "1.160",
		BasicAuthUser: "vsadmin",
		BasicAuthPassword: "secret",
		SSLVerify: false,
		Debug: false,
    		Timeout: 60 * time.Second,
	    },
	)
	//c.SetVserver("myvserver")
	volumeName := "vol_go_test01"
	response, _, err := c.VolumeContainerAPI(volumeName)
	if err != nil {
		fmt.Print(err.Error())
	} else {
		fmt.Printf("Containing Aggregate: %s\n", response.Results.ContainingAggregate)
	}
}
