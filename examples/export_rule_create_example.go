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
	options := &ontap.ExportRuleCreateOptions {
		    PolicyName:           "xpo_vol_go_test01",
		    AnonymousUserId:      "0",
            	    SuperUserSecurity:    &[]string{"any"},
            	    Protocol:             &[]string{"any"},
            	    IsAllowDevIsEnabled:  true,
            	    IsAllowSetUidEnabled: true,
            	    ClientMatch:          "192.168.1.5",
            	    RwRule:               &[]string{"any"},
            	    RoRule:               &[]string{"any"},
	}
	//c.SetVserver("myvserver")
	_, _, err := c.ExportRuleCreateAPI(options)
	if err != nil {
		fmt.Print(err.Error())
	} else {
		fmt.Println("Created export rule")
	}
}
