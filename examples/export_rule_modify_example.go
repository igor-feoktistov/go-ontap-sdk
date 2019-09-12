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
    options := &ontap.ExportRuleModifyOptions {
		PolicyName: "xpo_vol_go_test01",
		RuleIndex: 1,
                ClientMatch: "192.168.1.6",
    }
    //c.SetVserver("myvserver")
    _, _, err := c.ExportRuleModifyAPI(options)
    if err != nil {
	    fmt.Print(err.Error())
    } else {
	    fmt.Println("Modified export rule")
    }
}
