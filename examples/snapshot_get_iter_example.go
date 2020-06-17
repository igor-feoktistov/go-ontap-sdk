package main

import (
	"fmt"
	"time"
	"go-ontap-sdk/ontap"
)

func main() {
    var numRecords int
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
    options := &ontap.SnapshotGetOptions {
		    MaxRecords: 512,
		    },
    }
    //c.SetVserver("myvserver")
    response, err := c.SnapshotGetIterAPI(options)
    if err != nil {
	fmt.Print(err.Error())
    } else {
	    numRecords = 0
	    for _, responseSnapshot := range response {
		    numRecords += responseSnapshot.Results.NumRecords
		    for _, snapshot := range responseSnapshot.Results.AttributesList {
			    fmt.Printf("%s - %s\n", snapshot.Name, snapshot.VolumeProvenanceUuid)
		    }
	    }
	    fmt.Printf("Total Records: %d\n", numRecords)
    }
}
