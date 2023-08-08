package main

import (
	"fmt"
	"go-ontap-sdk/ontap"
	"time"
)

func main() {
	c := ontap.NewClient(
		"https://myvserver.example.com",
		&ontap.ClientOptions{
			Version:           "1.160",
			BasicAuthUser:     "vsadmin",
			BasicAuthPassword: "secret",
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
				fmt.Printf("Disk UID: %s\n", storageDisk.DiskUid)
				fmt.Printf("DiskRaid: %+v\n", storageDisk.DiskRaid)
			}
			fmt.Printf("Total Records: %d\n", response.Results.NumRecords)
		}
	}
}
