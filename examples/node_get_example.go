package main

import (
	"fmt"
	"go-ontap-sdk/ontap"
	"time"
)

func main() {
	c := ontap.NewClient(
		"https://mycluster.example.com",
		&ontap.ClientOptions{
			Version:           "1.160",
			BasicAuthUser:     "admin",
			BasicAuthPassword: "secret",
			SSLVerify:         false,
			Debug:             false,
			Timeout:           60 * time.Second,
		},
	)
	options := &ontap.SystemNodeGetOptions{
		MaxRecords: 1024,
	}
	response, _, err := c.SystemNodeGetAPI(options)
	if err != nil {
		fmt.Print(err)
	} else {
		if response.Results.NumRecords > 0 {
			for _, systemNode := range response.Results.SystemNodeAttributes.NodeDetailsInfo {
				fmt.Printf("Name: %s\n", systemNode.Node)
				fmt.Printf("SystemNodeAttributes: %+v\n", systemNode)
			}
			fmt.Printf("Total Records: %d\n", response.Results.NumRecords)
		}
	}
}
