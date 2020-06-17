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
	    BasicAuthPassword: "d0ntASK4IT",
	    SSLVerify: false,
	    Debug: false,
    	    Timeout: 60 * time.Second,
	},
    )
    options := &ontap.SnapshotRestoreVolumeOptions {
	    PreserveLunIds: true,
	    Volume: "vol_go_test01",
	    Snapshot: "vol_go_test01_snapshot01",
    }
    //c.SetVserver("myvserver")
    _, _, err := c.SnapshotRestoreVolumeAPI(options)
    if err != nil {
	    fmt.Print(err.Error())
    } else {
	    fmt.Printf("Restored volume from snapshot\n")
    }
}
