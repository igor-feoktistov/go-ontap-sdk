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
	volumeSize := "8g"
	response, _, err := c.VolumeSizeAPI(volumeName, volumeSize)
	if err != nil {
		fmt.Print(err.Error())
	} else {
		fmt.Printf("Volume Size: %s\n", response.Results.VolumeSize)
	}
}
