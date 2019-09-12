package ontap

import (
	"fmt"
	"encoding/xml"
	"net/http"
)

type IscsiInterfaceGetIter struct {
	Base
	Params struct {
		XMLName xml.Name
		*IscsiInterfaceGetOptions
	}
}

type IscsiInterfaceQuery struct {
	IscsiInterfaceListEntryInfo *IscsiInterfaceListEntryInfo `xml:"iscsi-interface-list-entry-info,omitempty"`
}

type IscsiInterfaceGetOptions struct {
	DesiredAttributes *IscsiInterfaceQuery `xml:"desired-attributes,omitempty"`
	MaxRecords        int                  `xml:"max-records,omitempty"`
	Query             *IscsiInterfaceQuery `xml:"query,omitempty"`
	Tag               string               `xml:"tag,omitempty"`
}

type IscsiInterfaceListEntryInfo struct {
	CurrentNode        string `xml:"current-node,omitempty"`
	CurrentPort        string `xml:"current-port,omitempty"`
	InterfaceName      string `xml:"interface-name,omitempty"`
	IpAddress          string `xml:"ip-address,omitempty"`
	IpPort             string `xml:"ip-port,omitempty"`
	IsInterfaceEnabled bool   `xml:"is-interface-enabled,omitempty"`
	RelativePortId     int    `xml:"relative-port-id,omitempty"`
	SendtargetsFqdn    string `xml:"sendtargets-fqdn,omitempty"`
	TpgroupName        string `xml:"tpgroup-name,omitempty"`
	TpgroupTag         int    `xml:"tpgroup-tag,omitempty"`
	Vserver            string `xml:"vserver,omitempty"`
}

type IscsiInterfaceGetResponse struct {
	XMLName xml.Name `xml:"netapp"`
	Results struct {
		ResultBase
		AttributesList struct {
			IscsiInterfaceAttributes []IscsiInterfaceListEntryInfo `xml:"iscsi-interface-list-entry-info"`
		} `xml:"attributes-list"`
		NextTag    string `xml:"next-tag"`
		NumRecords int    `xml:"num-records"`
	} `xml:"results"`
}

func (c *Client) IscsiInterfaceGetAPI(options *IscsiInterfaceGetOptions) (*IscsiInterfaceGetResponse, *http.Response, error) {
	if c.IscsiInterfaceGetIter == nil {
		c.IscsiInterfaceGetIter = &IscsiInterfaceGetIter {
			Base: Base {
				client:  c,
				XMLNs:   XMLNs,
				Version: c.options.Version,
			},
		}
		c.IscsiInterfaceGetIter.Params.XMLName = xml.Name{Local: "iscsi-interface-get-iter"}
	}
	c.IscsiInterfaceGetIter.Base.Name = c.vserver
	c.IscsiInterfaceGetIter.Params.IscsiInterfaceGetOptions = options
	r := IscsiInterfaceGetResponse{}
	res, err := c.IscsiInterfaceGetIter.get(c.IscsiInterfaceGetIter, &r)
	if err == nil && r.Results.Passed() == false {
	    err = fmt.Errorf("error(IscsiInterfaceGetAPI): %s", r.Results.Reason)
	}
	return &r, res, err
}

func (c *Client) IscsiInterfaceGetIterAPI(options *IscsiInterfaceGetOptions) (responseIscsiInterfaces []*IscsiInterfaceGetResponse, err error) {
	var nextTag string
	for {
		r, _, err := c.IscsiInterfaceGetAPI(options)
		if err == nil {
			nextTag = r.Results.NextTag
			responseIscsiInterfaces = append(responseIscsiInterfaces, r)
			if nextTag == "" {
				break
			}
			options.Tag = nextTag
		} else {
			break
		}
	}
	return
}
