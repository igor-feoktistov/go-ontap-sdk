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
	options := &ontap.VolumeCreateOptions {
			VolumeType: "rw",
			Volume: "vol_go_test01",
			JunctionPath: "/vol_go_test01",
			UnixPermissions: "0755",
			Size: "4g",
			ExportPolicy: "xpo_vol_go_test01",
			ContainingAggregateName: "myaggregate_01",
	}
	//c.SetVserver("myvserver")
	_, _, err := c.VolumeCreateAPI(options)
	if err != nil {
		fmt.Print(err.Error())
	} else {
		fmt.Println("Created volume")
	}
}
