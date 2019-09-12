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
	igroupType := "iscsi"
	osType := "linux"
	bindPortset := ""
	_, _, err := c.IgroupCreateAPI(igroupName, igroupType, osType, bindPortset)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Created initiator group")
	}
}
