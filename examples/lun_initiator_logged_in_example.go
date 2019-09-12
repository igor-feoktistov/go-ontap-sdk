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
	initiatorName := "iqn.1995-08.com.netapp:myhost01-dev"
	//c.SetVserver("myvserver")
	response, _, err := c.LunInitiatorLoggedInAPI(initiatorName)
	if err != nil {
		fmt.Print(err.Error())
	} else {
		if response.Results.IsLoggedIn {
			fmt.Println("Initiator is logged in")
		} else {
			fmt.Println("Initiator is not logged in")
		}
	}
}
