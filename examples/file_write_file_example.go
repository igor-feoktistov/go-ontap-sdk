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
	srcData := []byte("Test write via FileWriteFile.")
	dstData := make([]byte, hex.EncodedLen(len(srcData)))
	hex.Encode(dstData, srcData)
	options := &ontap.FileWriteFileOptions {
			Path: "/vol/vol_go_test01/TestFileWriteFile",
			Offset: 0,
			Data: string(dstData),
	}
	//c.SetVserver("myvserver")
	response, _, err := c.FileWriteFileAPI(options)
	if err != nil {
		fmt.Print(err.Error())
	} else {
		fmt.Printf("Wrote %d bytes to the file\n", response.Results.Length)
	}
}
