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
	response, _, err := c.IscsiServiceStatusAPI()
	if err != nil {
		fmt.Print(err)
	} else {
		if response.Results.IsAvailable {
			fmt.Println("iSCSI service is running")
		} else {
			fmt.Println("iSCSI service is not running")
		}
	}
}
