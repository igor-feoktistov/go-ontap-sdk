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
	options := &ontap.LunCreateFromFileOptions {
			FileName: "/vol/vol_go_test01/TestFileWriteFile",
			Path: "/vol/vol_go_test01/lun_go_test01",
			OsType: "linux",
	}
	//c.SetVserver("myvserver")
	response, _, err := c.LunCreateFromFileAPI(options)
	if err != nil {
		fmt.Print(err.Error())
	} else {
		fmt.Printf("Created LUN from file, actual size %d\n", response.Results.ActualSize)
	}
}
