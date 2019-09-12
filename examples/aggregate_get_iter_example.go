package main

import (
	"fmt"
	"time"
	"go-ontap-sdk/ontap"
)

func main() {
	var numRecords int
	c := ontap.NewClient(
		"https://mycluster.example.com",
		&ontap.ClientOptions {
		    Version: "1.160",
		    BasicAuthUser: "admin",
		    BasicAuthPassword: "secret",
		    SSLVerify: false,
		    Debug: false,
    		    Timeout: 60 * time.Second,
	    },
	)
	options := &ontap.AggregateGetOptions {
		    MaxRecords: 3,
	}
	response, err := c.AggregateGetIterAPI(options)
	if err != nil {
		fmt.Print(err.Error())
	} else {
		numRecords = 0
		for _, responseAggr := range response {
			numRecords += responseAggr.Results.NumRecords
    			for _, aggr := range responseAggr.Results.AggregateAttributes {
				fmt.Printf("%s\n", aggr.AggregateName)
				fmt.Printf("\tNode: %s\n", aggr.AggregateOwnershipAttributes.OwnerName)
				fmt.Printf("\tSize Total: %.2fGB\n", float64(aggr.AggregateSpaceAttributes.SizeTotal)/1024/1024/1024)
				fmt.Printf("\tSize Used: %.2fGB\n", float64(aggr.AggregateSpaceAttributes.SizeUsed)/1024/1024/1024)
				fmt.Printf("\tUsed Capacity: %s%%\n", aggr.AggregateSpaceAttributes.PercentUsedCapacity)
			}
		}
		fmt.Printf("Total Records: %d\n", numRecords)
	}
}
