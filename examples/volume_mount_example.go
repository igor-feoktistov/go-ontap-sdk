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
	//c.SetVserver("myvserver")
	volumeName := "vol_go_test01"
	junctionPath := "/vol_go_test01"
	_, _, err := c.VolumeMountAPI(volumeName, junctionPath, false)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Mounted volume")
	}
}
