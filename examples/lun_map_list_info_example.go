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
	lunPath := "/vol/vol_go_test01/lun_go_test01"
	response, _, err := c.LunMapListInfoAPI(lunPath)
	if err != nil {
		fmt.Print(err.Error())
	} else {
		fmt.Println("LUN mapped to the following initiator groups:")
		for _, igroup := range response.Results.InitiatorGroups.IgroupAttributes {
			fmt.Printf("\t%s\n", igroup.InitiatorGroupName)
			fmt.Printf("\t\tOS Type: %s\n", igroup.InitiatorGroupOsType)
			fmt.Println("\t\tInitators:")
    			for _, initiator := range igroup.Initiators {
				fmt.Printf("\t\t\tInitiator Name: %s\n", initiator.InitiatorName)
    			}
		}
	}
}
