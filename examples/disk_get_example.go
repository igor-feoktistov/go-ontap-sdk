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
			Version:           "1.32",
			BasicAuthUser:     "umonitor",
			BasicAuthPassword: "sxZeJs4n",
			SSLVerify:         false,
			Debug:             false,
			Timeout:           60 * time.Second,
		},
	)
	options := &ontap.StorageDiskGetOptions{
		MaxRecords: 1024,
	}
	response, _, err := c.StorageDiskGetAPI(options)
	if err != nil {
		fmt.Print(err)
	} else {
		if response.Results.NumRecords > 0 {
			for _, storageDisk := range response.Results.StorageDiskAttributes.DiskDetailsInfo {
				fmt.Printf("Name: %s\n", storageDisk.DiskName)
				fmt.Printf("storageDiskAttributes: %+v\n", storageDisk)
			}
			fmt.Printf("Total Records: %d\n", response.Results.NumRecords)
		}
	}
}
