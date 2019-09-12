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
	policyName := "xpo_vol_go_test01"
	_, _, err := c.ExportPolicyCreateAPI(policyName, false)
	if err != nil {
		fmt.Println(err)
	} else {
	    fmt.Println("Created export policy")
	}
}
