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
	//c.SetVserver("myvserver")
	response, _, err := c.LunGetAttributesAPI(lunPath, "")
	if err != nil {
		fmt.Print(err.Error())
	} else {
		for _, attribute := range response.Results.Attributes.LunAttributes {
			fmt.Printf("name=%s, value=\"%s\"\n", attribute.Name, attribute.Value)
		}
	}
}
