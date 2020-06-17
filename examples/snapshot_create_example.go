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
    options := &ontap.SnapshotCreateOptions {
		    Volume: "vol_go_test01",
		    Snapshot: "vol_go_test01_snapshot01",
    }
    //c.SetVserver("myvserver")
    _, _, err := c.SnapshotCreateAPI(options)
    if err != nil {
	    fmt.Print(err.Error())
    } else {
	    fmt.Printf("Created snapshot\n")
    }
}
