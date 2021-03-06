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
	lunPath := "/vol/vol_go_test02/lun_go_test02"
	attrName := "application.Name"
	attrValue := "My Application"
	//c.SetVserver("myvserver")
	_, _, err := c.LunSetAttributeAPI(lunPath, attrName, attrValue)
	if err != nil {
		fmt.Print(err.Error())
	} else {
		fmt.Println("Set LUN attribute")
	}
}
