package util

import (
	"fmt"
	"net"
	"github.com/igor-feoktistov/go-ontap-sdk/ontap"
)

func DiscoverIscsiLIFs(c *ontap.Client, lunPath string, initiatorSubnet string) (lifs []*ontap.NetInterfaceInfo, err error) {
	var lunInfo ontap.LunInfo
	options := &ontap.LunGetOptions {
			MaxRecords: 1,
			Query: &ontap.LunQuery {
				LunInfo: &ontap.LunInfo {
					Path: lunPath,
				},
			},
	}
	response, _, err := c.LunGetAPI(options)
	if err == nil {
		if response.Results.NumRecords > 0 {
			lunInfo = response.Results.AttributesList.LunAttributes[0]
			options := &ontap.NetInterfaceGetOptions {
					    MaxRecords: 64,
                			    Query: &ontap.NetInterfaceQuery {
                        			    NetInterfaceInfo: &ontap.NetInterfaceInfo {
                                			    DataProtocols: &[]string{"iscsi"},
                                			    Vserver: lunInfo.Vserver,
                        			    },
                			    },
			}
			response, _, err := c.NetInterfaceGetAPI(options)
			if err == nil {
    				for _, lifInfo := range response.Results.AttributesList.NetInterfaceAttributes {
    					if lifInfo.HomeNode == lunInfo.Node {
    						    if fmt.Sprintf("%s/%d", net.ParseIP(lifInfo.Address).Mask(net.CIDRMask(lifInfo.NetmaskLength, 32)), lifInfo.NetmaskLength) == initiatorSubnet {
    							    lif := lifInfo
    							    lifs = append(lifs, &lif)
    							    break
						    }
					}
				}
    				for _, lifInfo := range response.Results.AttributesList.NetInterfaceAttributes {
    					if lifInfo.HomeNode != lunInfo.Node {
    						if fmt.Sprintf("%s/%d", net.ParseIP(lifInfo.Address).Mask(net.CIDRMask(lifInfo.NetmaskLength, 32)), lifInfo.NetmaskLength) == initiatorSubnet {
    							lif := lifInfo
    							lifs = append(lifs, &lif)
    							break
						}
					}
				}
			}

		}
	}
	return lifs, err
}

func DiscoverNfsLIFs(c *ontap.Client, volume string) (lifs []*ontap.NetInterfaceInfo, err error) {
	var volumeInfo ontap.VolumeInfo
	options := &ontap.VolumeGetOptions {
			MaxRecords: 1,
			Query: &ontap.VolumeQuery {
                		VolumeInfo: &ontap.VolumeInfo {
					VolumeIDAttributes: &ontap.VolumeIDAttributes {
                        			Name: volume,
                        		},
                    		},
			},
	}
	response, _, err := c.VolumeGetAPI(options)
	if err == nil {
		if response.Results.NumRecords > 0 {
			volumeInfo = response.Results.AttributesList[0]
			options := &ontap.NetInterfaceGetOptions {
					    MaxRecords: 64,
                			    Query: &ontap.NetInterfaceQuery {
                        			    NetInterfaceInfo: &ontap.NetInterfaceInfo {
                                			    DataProtocols: &[]string{"nfs"},
                                			    Vserver: volumeInfo.VolumeIDAttributes.OwningVserverName,
                        			    },
                			    },
			}
			response, _, err := c.NetInterfaceGetAPI(options)
			if err == nil {
    				for _, lifInfo := range response.Results.AttributesList.NetInterfaceAttributes {
    					if lifInfo.HomeNode == volumeInfo.VolumeIDAttributes.Node {
    						lif := lifInfo
    						lifs = append(lifs, &lif)
					}
				}
    				for _, lifInfo := range response.Results.AttributesList.NetInterfaceAttributes {
    					if lifInfo.HomeNode != volumeInfo.VolumeIDAttributes.Node {
    						lif := lifInfo
    						lifs = append(lifs, &lif)
					}
				}
			}
		}
	}
	return lifs, err
}
