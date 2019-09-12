package main

import (
	"fmt"
	"time"
	"encoding/hex"
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
	options := &ontap.FileReadFileOptions {
			Path: "/vol/vol_go_test01/TestFileWriteFile",
			Offset: 0,
			Length: 2048,
	}
	//c.SetVserver("myvserver")
	response, _, err := c.FileReadFileAPI(options)
	if err != nil {
		fmt.Print(err.Error())
	} else {
		srcData := []byte(response.Results.Data)
		dstData := make([]byte, hex.DecodedLen(len(srcData)))
		hex.Decode(dstData, srcData)
		fmt.Printf("%s\n", dstData)
	}
}
