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
	options := &ontap.LunResizeOptions {
			Path: "/vol/vol_go_test01/lun_go_test01",
			Size: 5368709120,
	}
	//c.SetVserver("myvserver")
	response, _, err := c.LunResizeAPI(options)
	if err != nil {
		fmt.Print(err.Error())
	} else {
		fmt.Printf("Resized LUN, actual size %d\n", response.Results.ActualSize)
	}
}
