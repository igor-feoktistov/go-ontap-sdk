package main

import (
	"fmt"
	"time"
	"strings"
	"net"
	"go-ontap-sdk/ontap"
	"go-ontap-sdk/util"
)

func main() {
	c := ontap.NewClient(
		"https://mycluster.example.com",
		&ontap.ClientOptions {
		    Version: "1.160",
		    BasicAuthUser: "vsadmin",
		    BasicAuthPassword: "secret",
		    SSLVerify: false,
		    Debug: false,
    		    Timeout: 60 * time.Second,
		},
	)
	lunPath := "/vol/vol_go_test01/lun_go_test01"
	initiatorSubnet := "192.168.1.0/24"
	lifs, err := util.DiscoverIscsiLIFs(c, lunPath, initiatorSubnet)
	if err != nil {
		fmt.Println(err)
	} else {
		for _, lif := range lifs {
			fmt.Printf("%s\n", lif.InterfaceName)
			fmt.Printf("\tAddress: %s\n", lif.Address)
			fmt.Printf("\tNetmask: %s\n", lif.Netmask)
			fmt.Printf("\tNetwork: %s/%d\n", net.ParseIP(lif.Address).Mask(net.CIDRMask(lif.NetmaskLength, 32)), lif.NetmaskLength)
			fmt.Printf("\tProtocols: %s\n", strings.Join(*lif.DataProtocols, ","))
			fmt.Printf("\tHome Node: %s\n", lif.HomeNode)
			fmt.Printf("\tHome Port: %s\n", lif.HomePort)
		}
	}
}
