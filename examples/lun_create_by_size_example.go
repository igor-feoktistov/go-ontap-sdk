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
	options := &ontap.LunCreateBySizeOptions {
			Path: "/vol/vol_go_test01/lun_go_test01",
			Size: 1073741824,
			OsType: "linux",
	}
	//c.SetVserver("myvserver")
	response, _, err := c.LunCreateBySizeAPI(options)
	if err != nil {
		fmt.Print(err.Error())
	} else {
		fmt.Printf("Created LUN, actual size %d\n", response.Results.ActualSize)
	}
}
