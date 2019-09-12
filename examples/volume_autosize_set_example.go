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
	options := &ontap.VolumeAutosizeOptions {
			IsEnabled: true,
			IncrementSize: "2g",
			MaximumSize: "8g",
			GrowThresholdPercent: 85,
			Volume: "vol_go_test01",
	}
	//c.SetVserver("myvserver")
	_, _, err := c.VolumeAutosizeSetAPI(options)
	if err != nil {
		fmt.Print(err.Error())
	} else {
		fmt.Println("Set Autosize options for volume")
	}
}
