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
	igroupName := "myhost01_dev_igroup01"
	initiatorName := "iqn.2005-02.com.open-iscsi:myhist01-dev"
	force := true
	_, _, err := c.IgroupRemoveAPI(igroupName, initiatorName, force)
	if err != nil {
		fmt.Println(err)
	} else {
	    fmt.Println("Removed initiator name")
	}
}
