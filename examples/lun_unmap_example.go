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
	options := &ontap.LunUnmapOptions {
			InitiatorGroup: "myhost01_dev_igroup01",
			Path: "/vol/vol_go_test01/lun_go_test01",
	}
	//c.SetVserver("myvserver")
	_, _, err := c.LunUnmapAPI(options)
	if err != nil {
		fmt.Print(err.Error())
	} else {
		fmt.Println("Unmapped LUN")
	}
}
